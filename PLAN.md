# CLI Word Frequency Analyzer - Implementation Plan

## Overview
Build a CLI that takes a URL, fetches content, and displays the top 10 most frequent words.

## Steps

### Step 1: URL Fetching
Create a function that takes a URL and returns the raw HTML content.

**Test:** Pass a known URL (e.g., `https://example.com`) and verify HTML is returned.

### Step 2: Text Extraction
Parse the HTML and extract plain text, stripping tags and scripts.

**Test:** Pass sample HTML like `<p>Hello world</p>` and verify output is `Hello world`.

### Step 3: Word Counting
Tokenize the text into words and count frequencies, ignoring case and punctuation.

**Test:** Pass `"the cat and the dog"` and verify `{"the": 2, "cat": 1, "and": 1, "dog": 1}`.

### Step 4: CLI Interface
Wire everything together with argument parsing and formatted output of top 10 words.

**Test:** Run `./cli https://example.com` and verify formatted output appears.
