package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

type wordCount struct {
	word  string
	count int
}

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func extractText(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	var textParts []string
	var extract func(*html.Node)
	extract = func(n *html.Node) {
		// Skip script and style elements
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
			return
		}

		if n.Type == html.TextNode {
			text := strings.TrimSpace(n.Data)
			if text != "" {
				textParts = append(textParts, text)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}

	extract(doc)
	return strings.Join(textParts, " "), nil
}

var nonAlpha = regexp.MustCompile(`[^a-zA-Z]+`)

func countWords(text string) map[string]int {
	counts := make(map[string]int)

	// Normalize: lowercase and replace non-alpha with spaces
	normalized := nonAlpha.ReplaceAllString(strings.ToLower(text), " ")

	for _, word := range strings.Fields(normalized) {
		if word != "" {
			counts[word]++
		}
	}

	return counts
}

func topN(counts map[string]int, n int) []wordCount {
	words := make([]wordCount, 0, len(counts))
	for word, count := range counts {
		words = append(words, wordCount{word, count})
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].count != words[j].count {
			return words[i].count > words[j].count
		}
		return words[i].word < words[j].word // alphabetical tiebreaker
	})

	if n > len(words) {
		n = len(words)
	}
	return words[:n]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: wordfreq <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	content, err := fetchURL(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching URL: %v\n", err)
		os.Exit(1)
	}

	text, err := extractText(content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting text: %v\n", err)
		os.Exit(1)
	}

	counts := countWords(text)
	top := topN(counts, 10)

	fmt.Println("Top 10 words:")
	for i, wc := range top {
		fmt.Printf("%2d. %-15s %d\n", i+1, wc.word, wc.count)
	}
}
