# Technical Context: Degen's Digest

**Core Language:**
-   N/A (Process driven by Cline and MCP tools, no custom script required).

**Key Dependencies/Tools:**
-   **Reddit MCP Server:** Accessed via `use_mcp_tool`. Responsible for fetching Reddit data:
    -   Fetching hot threads (`fetch_reddit_hot_threads`).
    -   Fetching post content and comments (`fetch_reddit_post_content`).
-   **Cline's NLP Capabilities:** Used for parsing comment text and extracting structured data.
-   **Standard File Tools:** `write_to_file` used for generating the final CSV report.
-   **Target Website:**
    -   `https://www.reddit.com/r/sportsbook`: The source of the data.

**Development Environment:**
-   **OS:** macOS (as per SYSTEM INFORMATION).
-   **Shell:** zsh (as per SYSTEM INFORMATION).
-   **MCP Servers:** Reddit server (`reddit`) must be running and accessible.

**Technical Constraints:**
-   **Reddit MCP Tool Limitations:** The accuracy and completeness of data depend on the capabilities and reliability of the `fetch_reddit_hot_threads` and `fetch_reddit_post_content` tools.
-   **NLP Parsing Robustness:** While generally robust, NLP might misinterpret highly unconventional comment formats or fail to extract all desired fields if they are presented ambiguously.
-   **Reddit Rate Limiting:** The Reddit MCP server might be subject to Reddit's API rate limits.

**Setup:**
1.  Ensure the Reddit MCP server is running.
2.  Instruct Cline to perform the POTD scraping task.
