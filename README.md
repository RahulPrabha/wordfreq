# wordfreq

A CLI tool that fetches a URL and displays the top 10 most frequent words.

## Installation

```bash
go install github.com/yourusername/wordfreq@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/wordfreq.git
cd wordfreq
go build -o wordfreq .
```

## Usage

```bash
wordfreq <url>
```

Example:

```bash
$ wordfreq https://example.com
Top 10 words:
 1. domain          3
 2. example         2
 3. in              2
 4. use             2
 5. avoid           1
 6. documentation   1
 7. examples        1
 8. for             1
 9. is              1
10. learn           1
```

## How it works

1. Fetches HTML content from the URL
2. Extracts plain text, stripping scripts and styles
3. Tokenizes into words, normalizing to lowercase and removing punctuation
4. Counts frequencies and displays top 10 (alphabetical tiebreaker)

## Running tests

```bash
go test -v
```
