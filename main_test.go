package main

import (
	"strings"
	"testing"
)

func TestFetchURL(t *testing.T) {
	content, err := fetchURL("https://example.com")
	if err != nil {
		t.Fatalf("fetchURL failed: %v", err)
	}

	if len(content) == 0 {
		t.Fatal("fetchURL returned empty content")
	}

	if !strings.Contains(content, "Example Domain") {
		t.Error("expected content to contain 'Example Domain'")
	}

	if !strings.Contains(content, "<html") {
		t.Error("expected content to contain HTML tags")
	}
}

func TestFetchURL_InvalidURL(t *testing.T) {
	_, err := fetchURL("https://this-domain-does-not-exist-12345.com")
	if err == nil {
		t.Error("expected error for invalid URL")
	}
}

func TestExtractText(t *testing.T) {
	html := "<p>Hello world</p>"
	text, err := extractText(html)
	if err != nil {
		t.Fatalf("extractText failed: %v", err)
	}

	if text != "Hello world" {
		t.Errorf("expected 'Hello world', got '%s'", text)
	}
}

func TestExtractText_SkipsScriptAndStyle(t *testing.T) {
	html := `<html>
		<head><style>body { color: red; }</style></head>
		<body>
			<script>alert('hi');</script>
			<p>Visible text</p>
		</body>
	</html>`

	text, err := extractText(html)
	if err != nil {
		t.Fatalf("extractText failed: %v", err)
	}

	if !strings.Contains(text, "Visible text") {
		t.Error("expected text to contain 'Visible text'")
	}

	if strings.Contains(text, "alert") {
		t.Error("expected script content to be stripped")
	}

	if strings.Contains(text, "color") {
		t.Error("expected style content to be stripped")
	}
}
