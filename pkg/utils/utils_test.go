package utils

import (
	"testing"
)

func TestRewritePixeldrainURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "https://pixeldrain.com/u/abc12345",
			expected: "https://pixeldrain.com/api/file/abc12345",
		},
		{
			input:    "https://pixeldrain.com/api/file/abc12345",
			expected: "https://pixeldrain.com/api/file/abc12345",
		},
		{
			input:    "https://google.com",
			expected: "https://google.com",
		},
	}

	for _, tt := range tests {
		got := RewritePixeldrainURL(tt.input)
		if got != tt.expected {
			t.Errorf("RewritePixeldrainURL(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}

func TestIsValidURL(t *testing.T) {
	if !IsValidURL("https://example.com") {
		t.Errorf("Expected true for valid URL")
	}
	if IsValidURL("not-a-url") {
		t.Errorf("Expected false for invalid URL")
	}
}

func TestGetHealthStatus(t *testing.T) {
	status := GetHealthStatus()
	if status["status"] != "ok" {
		t.Errorf("Expected status ok, got %s", status["status"])
	}
}
