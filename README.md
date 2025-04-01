# Degen's Digest - r/sportsbook POTD Scraper

This project contains a Go application designed to scrape the daily "Pick of the Day" (POTD) thread from the r/sportsbook subreddit. It extracts key information from user comments and compiles it into a daily CSV report.

## Purpose

The goal is to automate the process of collecting daily sports picks shared by users in the POTD thread, saving time and providing a structured dataset for review.

## How it Works

1.  The script uses the Playwright MCP server to launch a browser instance.
2.  It navigates to `https://www.reddit.com/r/sportsbook`.
3.  It searches for the POTD thread corresponding to the current date (e.g., by matching the title pattern "Pick of the Day - Month Day, Year").
4.  It navigates to the identified POTD thread.
5.  It fetches the HTML content of the comments section.
6.  The Go application parses this HTML, looking for comments that appear to contain sports picks. It attempts to extract:
    *   Username (if identifiable)
    *   User's POTD Record (W-L-P format)
    *   Today's Pick
    *   Sport
    *   Game Details
    *   Reasoning/Write-up
7.  The extracted data is written to a CSV file named `potd_picks_YYYY-MM-DD.csv` in the root directory of this project.

## Usage

1.  **Prerequisites:**
    *   Go (version specified in `go.mod`, e.g., 1.23.6 or later) must be installed and configured.
    *   The Playwright MCP server (`github.com/executeautomation/mcp-playwright`) must be running and accessible by the system.
2.  **Run the script:**
    Open a terminal in the project's root directory (`/Users/jasonlane/Documents/code/degens-digest`) and execute:
    ```bash
    go run main.go
    ```
3.  **Output:**
    A CSV file named `potd_picks_YYYY-MM-DD.csv` (e.g., `potd_picks_2025-04-01.csv`) will be created or overwritten in the project root directory upon successful execution.

## Dependencies

-   Go Standard Library
-   Playwright MCP Server (External Tool)

## Limitations

-   **Web Scraping Fragility:** This script relies on the specific HTML structure and CSS classes of Reddit. Changes to Reddit's website design *will* likely break the scraper, requiring updates to the parsing logic (specifically the regular expressions and potentially the Playwright selectors) in `main.go`.
-   **Inconsistent Comment Formatting:** Users format their picks in many different ways. The regular expressions used for parsing are designed to be somewhat flexible but may fail to extract data correctly (or at all) from comments that deviate significantly from common patterns.
-   **POTD Thread Identification:** The script assumes a consistent title format for the daily POTD thread. If the moderators change the naming convention, the script may fail to find the correct thread.
-   **Reddit Anti-Scraping Measures:** Running the script too frequently or aggressively might trigger rate limiting or temporary blocks from Reddit.

## Development Notes

-   The core parsing logic is within the `parsePicksFromHTML` function in `main.go`. This is the most likely area needing refinement as edge cases or new comment formats are encountered.
-   The actual interaction with Playwright (navigation, fetching HTML) is intended to be handled by invoking the `use_mcp_tool` calls *before* running `go run main.go`, or by modifying `main.go` to accept the HTML content as input (e.g., via stdin or a temporary file). The current `main.go` uses placeholder functions.
-   The `.clinerules` file and `memory-bank/` directory are used for project context and learning during development with Cline.
