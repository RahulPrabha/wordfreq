package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

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
	fmt.Printf("Word counts: %v\n", counts)
}
