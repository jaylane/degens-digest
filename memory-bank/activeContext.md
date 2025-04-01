# Active Context: Degen's Digest (Initial Setup)

**Current Focus:** Initial project setup and development of the core scraping and parsing logic in Go.

**Recent Changes:**
-   Initialized the project directory (`/Users/jasonlane/Documents/code/degens-digest`).
-   Created the `memory-bank/` directory.
-   Created core Memory Bank documents:
    -   `projectbrief.md`
    -   `productContext.md`
    -   `systemPatterns.md`
    -   `techContext.md`

**Next Steps:**
1.  Create `memory-bank/progress.md`.
2.  Create the `.clinerules` file.
3.  Implement the main Go application (`main.go`), including:
    -   Playwright interaction logic (navigation, HTML fetching).
    -   HTML parsing logic using regex.
    -   CSV generation logic.
4.  Create the `README.md` file.
5.  Test the script by running it.

**Active Decisions & Considerations:**
-   **Parsing Strategy:** Shifted from regex parsing within Go. The Go script will now extract raw comment text. Detailed parsing (Record, Pick, Sport, Game, Reasoning) will be handled by Cline's NLP capabilities after the Go script generates a raw comment CSV.
-   **Error Handling:** Need to implement robust error checks at each stage (Playwright actions, Go text extraction, file I/O, NLP processing).
-   **POTD Thread Identification:** Will use date-based title matching (e.g., "Pick of the Day - April 1, 2025") via Playwright to find the correct thread.
-   **Data Flow:** Playwright (HTML) -> Go (Raw Comment Text Extraction -> `raw_...csv`) -> Cline NLP (Detailed Parsing) -> Final `potd_picks_...csv`.
