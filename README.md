# Degen's Digest - r/sportsbook POTD Scraper

This project automates the collection of daily "Pick of the Day" (POTD) entries from the r/sportsbook subreddit using the Reddit MCP server and NLP processing. It extracts key information from top-level user comments and compiles it into a daily CSV report.

## Purpose

The goal is to automate the process of collecting daily sports picks shared by users in the POTD thread, saving time and providing a structured dataset for review.

## How it Works

1.  **Fetch Hot Threads:** Uses the `fetch_reddit_hot_threads` tool (Reddit MCP server) to get recent posts from r/sportsbook.
2.  **Identify POTD Post:** Analyzes thread titles to find the POTD post for the current date and extracts its ID.
3.  **Fetch Comments:** Uses the `fetch_reddit_post_content` tool (Reddit MCP server) with the post ID to retrieve top-level comments (depth=1).
4.  **NLP Processing:** Analyzes the text of each top-level comment to:
    *   Filter out non-pick comments (e.g., replies, general discussion).
    *   Extract structured data for valid picks: Username, Record, Today's Pick, Sport/League, Game Details, Reasoning.
5.  **Generate Report:** Writes the extracted structured data to a CSV file named `potd_picks_YYYY-MM-DD.csv` in the project root directory.

## Usage

This process is intended to be run via prompts to Cline (or a similar AI assistant with access to the required tools).

1.  **Prerequisites:**
    *   The Reddit MCP server (`reddit`) must be running and accessible.
2.  **Execution:**
    Instruct Cline to perform the POTD scraping task for the desired date. Cline will execute the steps outlined in "How it Works".
3.  **Output:**
    A CSV file named `potd_picks_YYYY-MM-DD.csv` (e.g., `potd_picks_2025-04-01.csv`) containing the structured pick data will be created or overwritten in the project root directory (`/Users/jasonlane/Documents/code/degens-digest`).

## Dependencies

-   Reddit MCP Server (External Tool)
-   Cline's NLP Capabilities

## Limitations

-   **Reddit API/Tool Changes:** If the Reddit MCP server's functionality changes or Reddit alters its API/structure in a way the tool cannot handle, the process may fail.
-   **POTD Thread Title Format:** Relies on a predictable title format for the daily POTD thread (e.g., "Pick of the Day - M/D/YY (Day)"). Changes to this format will break thread identification.
-   **Inconsistent Comment Formatting:** While NLP is used, highly unusual or ambiguous comment formatting might still lead to incorrect data extraction or missed picks.
-   **Reddit Rate Limiting:** The Reddit MCP server might be subject to Reddit's API rate limits.

## Development Notes

-   This workflow replaces the previous Go script/Playwright/manual HTML approach.
-   The core data extraction now relies on NLP analysis performed by Cline after fetching comments via the Reddit MCP tool.
-   The `.clinerules` file and `memory-bank/` directory are used for project context and learning during development with Cline.
