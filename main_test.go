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
