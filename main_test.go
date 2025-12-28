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

func TestCountWords(t *testing.T) {
	counts := countWords("the cat and the dog")

	expected := map[string]int{
		"the": 2,
		"cat": 1,
		"and": 1,
		"dog": 1,
	}

	for word, count := range expected {
		if counts[word] != count {
			t.Errorf("expected %s=%d, got %d", word, count, counts[word])
		}
	}
}

func TestCountWords_IgnoresCaseAndPunctuation(t *testing.T) {
	counts := countWords("Hello, HELLO! hello.")

	if counts["hello"] != 3 {
		t.Errorf("expected hello=3, got %d", counts["hello"])
	}

	// Should not have entries with punctuation
	for word := range counts {
		if strings.ContainsAny(word, ",.!") {
			t.Errorf("unexpected punctuation in word: %s", word)
		}
	}
}

func TestCountWords_EmptyString(t *testing.T) {
	counts := countWords("")
	if len(counts) != 0 {
		t.Errorf("expected empty map, got %v", counts)
	}
}

func TestCountWords_OnlyPunctuation(t *testing.T) {
	counts := countWords("!@#$%^&*()")
	if len(counts) != 0 {
		t.Errorf("expected empty map, got %v", counts)
	}
}

func TestCountWords_OnlyWhitespace(t *testing.T) {
	counts := countWords("   \t\n   ")
	if len(counts) != 0 {
		t.Errorf("expected empty map, got %v", counts)
	}
}

func TestCountWords_SingleWord(t *testing.T) {
	counts := countWords("hello")
	if counts["hello"] != 1 || len(counts) != 1 {
		t.Errorf("expected {hello:1}, got %v", counts)
	}
}

func TestCountWords_NumbersIgnored(t *testing.T) {
	counts := countWords("test123 456 test")
	// "test123" becomes "test" after stripping non-alpha, "456" becomes empty
	if counts["test"] != 2 {
		t.Errorf("expected test=2, got %v", counts)
	}
}
