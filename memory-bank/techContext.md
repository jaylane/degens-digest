# Technical Context: Degen's Digest

**Core Language:**
-   **Go:** Version specified in `go.mod` (or latest stable if not yet defined). Standard library features will be prioritized.

**Key Dependencies/Tools:**
-   **Go Standard Library:**
    -   `fmt`: For printing output/logs.
    -   `os`: For file system operations (creating/writing CSV).
    -   `time`: For handling dates (generating filename, finding today's thread).
    -   `encoding/csv`: For writing data to the CSV file.
    -   `strings`: For text manipulation during parsing.
    -   `regexp`: For pattern matching to extract data from HTML.
    -   `net/http`: Potentially for direct API calls if Playwright proves insufficient or too slow, though Playwright is the primary plan.
-   **Playwright MCP Server:** Accessed via `use_mcp_tool`. Responsible for all browser automation tasks:
    -   Navigating to URLs (`playwright_navigate`).
    -   Clicking elements (`playwright_click`).
    -   Fetching rendered HTML (`playwright_get_visible_html`).
    -   Executing JavaScript if needed (`playwright_evaluate`).
    -   Closing the browser session (`playwright_close`).
-   **Target Website:**
    -   `https://www.reddit.com/r/sportsbook`: The source of the data.

**Development Environment:**
-   **OS:** macOS (as per SYSTEM INFORMATION).
-   **Shell:** zsh (as per SYSTEM INFORMATION).
-   **Go Installation:** Assumed to be present and configured correctly on the user's system.
-   **MCP Servers:** Playwright server (`github.com/executeautomation/mcp-playwright`) must be running and accessible.

**Technical Constraints:**
-   **Web Scraping Limitations:** Subject to Reddit's terms of service, rate limits, and potential anti-bot measures. The script must be designed to be relatively "polite" (e.g., avoid extremely rapid requests).
-   **HTML Parsing Robustness:** Using regex for HTML parsing is inherently less robust than dedicated HTML parsing libraries. Complex or malformed HTML might cause parsing failures.
-   **No External Go Libraries (Initial Plan):** Sticking to the standard library initially. If regex parsing becomes unmanageable, libraries like `golang.org/x/net/html` or `github.com/PuerkitoBio/goquery` might be considered, requiring updates to `go.mod`.

**Setup:**
1.  Ensure Go is installed.
2.  Ensure the Playwright MCP server is running.
3.  Run the script using `go run main.go` from the project root (`/Users/jasonlane/Documents/code/degens-digest`).
