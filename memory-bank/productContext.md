# Product Context: Degen's Digest

**Problem:** Manually tracking daily sports picks posted by users in the r/sportsbook "Pick of the Day" (POTD) thread is time-consuming and tedious. Users share valuable insights, including their track records, reasoning, and specific game picks, but this information is scattered across numerous comments with inconsistent formatting.

**Solution:** Develop an automated tool ("Degen's Digest") that scrapes the daily POTD thread, extracts key information from user comments, and aggregates it into a structured CSV format.

**Why:**
-   **Efficiency:** Saves the user significant time compared to manual collection.
-   **Data Aggregation:** Provides a consolidated view of daily picks for easier analysis and tracking.
-   **Trend Spotting:** Over time, the generated CSVs could potentially be used to identify trends or successful predictors within the subreddit community (though this analysis is outside the scope of the initial tool).

**How it Should Work:**
1.  The user runs the Go script.
2.  The script automatically navigates to r/sportsbook using Playwright.
3.  It identifies and navigates to the POTD thread for the current date.
4.  It fetches the HTML content of the thread's comments.
5.  It parses the comments, attempting to extract predefined data points (Username, Record, Today's Pick, Sport, Game Details, Reasoning).
6.  It generates a CSV file named `potd_picks_YYYY-MM-DD.csv` in the project root, containing the extracted data.

**User Experience Goals:**
-   **Simple Execution:** The script should be runnable with a single command (`go run main.go`).
-   **Clear Output:** The CSV report should be well-formatted with clear headers.
-   **Reliable (within limits):** The script should function correctly on most days, acknowledging the inherent fragility of web scraping. Error handling should inform the user of issues (e.g., inability to find the thread, parsing errors).
