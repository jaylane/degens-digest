package main

import (
	"encoding/csv"
	"fmt"
	"html" // For unescaping HTML entities
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	// NOTE: This script now expects HTML content to be provided via a local file,
	// as direct scraping via Playwright was blocked.
)

// CommentData holds the raw text content and username for a single comment.
type CommentData struct {
	Username    string
	CommentText string
}

func main() {
	log.Println("Starting Degen's Digest raw comment extractor...")

	// 1. Determine input HTML filename
	dateStr := time.Now().Format("2006-01-02")
	htmlFilename := fmt.Sprintf("potd_comments_%s.html", dateStr)
	if len(os.Args) > 1 {
		htmlFilename = os.Args[1]
		log.Printf("Using provided HTML file: %s\n", htmlFilename)
	} else {
		log.Printf("Using default HTML file based on date: %s\n", htmlFilename)
	}

	// 2. Read HTML content from the file
	htmlBytes, err := os.ReadFile(htmlFilename)
	if err != nil {
		log.Fatalf("Failed to read HTML file %s: %v\nMake sure you have saved the POTD thread HTML to this file.", htmlFilename, err)
		return // Added return for clarity
	}
	htmlContent := string(htmlBytes)
	log.Printf("Successfully read %d bytes from %s\n", len(htmlContent), htmlFilename)

	// 3. Parse the HTML content to extract raw comments
	comments := parseCommentsFromHTML(htmlContent)
	if len(comments) == 0 {
		log.Println("No comments found or parsed from the HTML content.")
	} else {
		log.Printf("Parsed %d potential comments.\n", len(comments))
	}

	// 4. Write raw comments to CSV (Changed step number for clarity)
	err = writeCommentsToCSV(comments) // Use = instead of := as err is already declared
	if err != nil {
		log.Fatalf("Failed to write comments to CSV: %v\n", err)
	}

	log.Println("Raw comment extraction process finished successfully.")
}

// getPlaceholderHTML is no longer used, replaced by file reading logic.
/*
func getPlaceholderHTML() string {
	// This is dummy HTML representing a simplified comment structure.
	// The actual HTML from Reddit will be much more complex.
	return `
<div>
  <div class="comment">
    <p>Username: User1</p>
    <p>Record: 10-5-1</p>
    <p>Yesterday's Pick: Lost</p>
    <p>Today's Pick: Team A -3.5</p>
    <p>Sport: NBA</p>
    <p>Game: Team A vs Team B</p>
    <p>Reasoning: Team A is hot right now...</p>
  </div>
  <div class="comment">
    <p>Record: 5-2-0. Today's Pick: Over 2.5 Goals. Sport: Soccer. Game: Club X vs Club Y. They always score a lot.</p>
  </div>
  <div class="comment non-pick">
    <p>Just here to tail.</p>
  </div>
</div>
`
}
*/

// Simple regex to strip HTML tags. More robust parsing might be needed.
var stripTagsRegex = regexp.MustCompile("<[^>]*>")

// parseCommentsFromHTML extracts raw comment text from the HTML content read from the file,
// using the refined structure identified from the saved HTML.
func parseCommentsFromHTML(htmlContent string) []CommentData {
	var comments []CommentData

	// More flexible Regex to find TOP-LEVEL comment blocks (depth="0") AND capture author and inner content,
	// allowing attributes to be in different orders.
	// Group 1: author if it comes before depth
	// Group 2: author if it comes after depth
	// Group 3: inner content of the shreddit-comment tag
	commentBlockRegex := regexp.MustCompile(`(?s)<shreddit-comment(?:[^>]*\bauthor="([^"]+)"[^>]*\bdepth="0"|[^>]*\bdepth="0"[^>]*\bauthor="([^"]+)")[^>]*>(.*?)</shreddit-comment>`)

	// Regex to find the comment text div using the slot="comment" attribute (applied to inner content)
	commentTextDivRegex := regexp.MustCompile(`(?s)<div[^>]*slot="comment"[^>]*>(.*?)</div>`)

	commentMatches := commentBlockRegex.FindAllStringSubmatch(htmlContent, -1)
	log.Printf("Found %d potential top-level <shreddit-comment> blocks (depth=0) with author attribute (flexible order).\n", len(commentMatches))

	for i, match := range commentMatches {
		if len(match) < 4 { // Expecting full match, author1|author2, inner content
			log.Printf("Skipping match %d due to unexpected number of capture groups (%d)\n", i+1, len(match))
			continue
		}

		// Determine username from the correct capture group
		username := ""
		if match[1] != "" {
			username = match[1] // Author came before depth
		} else if match[2] != "" {
			username = match[2] // Author came after depth
		} else {
			log.Printf("Could not extract username from match %d despite matching block structure.\n", i+1)
			username = "Unknown" // Fallback
		}

		commentInnerContent := match[3] // The content inside the <shreddit-comment> tag

		// Extract comment text div content from the inner content
		plainText := ""
		if textMatch := commentTextDivRegex.FindStringSubmatch(commentInnerContent); len(textMatch) > 1 {
			commentTextHTML := textMatch[1]
			// Extract raw text by stripping HTML tags from the comment div's content
			plainText = stripTagsRegex.ReplaceAllString(commentTextHTML, "")
			// Unescape HTML entities like &, <, etc.
			plainText = html.UnescapeString(plainText)
			// Trim whitespace
			plainText = strings.TrimSpace(plainText)
		} else {
			log.Printf("Could not find comment text div (slot=\"comment\") pattern in comment block %d for user %s\n", i+1, username)
		}

		// Only add if we found actual text content
		if plainText != "" {
			comments = append(comments, CommentData{
				Username:    username, // Use username from author attribute
				CommentText: plainText,
			})
		} else {
			log.Printf("Skipping comment block %d (user: %s), no text content found after processing.\n", i+1, username)
		}
	}

	return comments
}

// writeCommentsToCSV writes the extracted raw comment data to a CSV file.
func writeCommentsToCSV(comments []CommentData) error {
	dateStr := time.Now().Format("2006-01-02")                   // YYYY-MM-DD format
	filename := fmt.Sprintf("raw_potd_comments_%s.csv", dateStr) // Changed filename

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Username", "Comment Text"} // Updated header
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header to csv: %w", err)
	}

	// Write comment data
	for _, comment := range comments {
		row := []string{
			comment.Username,
			comment.CommentText,
		}
		if err := writer.Write(row); err != nil {
			// Log error but continue trying to write other rows
			log.Printf("Error writing row to csv for user '%s': %v\n", comment.Username, err)
		}
	}

	log.Printf("Successfully wrote %d raw comments to %s\n", len(comments), filename)
	return writer.Error() // Check for any final error during flush
}

// TODO:
// 1. **Manual Step:** User needs to save the full HTML source of the daily POTD thread to `potd_comments_YYYY-MM-DD.html` (or provide filename as arg). (DONE for 2025-04-01)
// 2. **Refine Parsing:** Comment regex updated to be more flexible with attribute order for depth="0" and author. Check logs.
// 3. **Refine HTML Stripping:** Current stripping is basic; might need adjustment for complex formatting within comments (e.g., nested blockquotes, code blocks).
// 4. **Error Handling:** Basic logging added for missing patterns. Could be enhanced.
// 5. **Post-processing:** After this script runs, the resulting CSV (`raw_potd_comments_YYYY-MM-DD.csv`) needs to be processed by Cline's NLP capabilities to extract the detailed pick information (Record, Pick, Sport, Game, Reasoning) into the final `potd_picks_YYYY-MM-DD.csv` report, with stricter filtering for actual picks.
