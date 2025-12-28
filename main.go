package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

	fmt.Printf("Fetched %d bytes\n", len(content))
}
