package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// RewritePixeldrainURL converts a Pixeldrain web viewer URL (e.g. pixeldrain.com/u/ID)
// to a direct file download URL (e.g. pixeldrain.com/api/file/ID).
func RewritePixeldrainURL(rawURL string) string {
	if strings.Contains(rawURL, "pixeldrain.com/u/") {
		return strings.ReplaceAll(rawURL, "pixeldrain.com/u/", "pixeldrain.com/api/file/")
	}
	return rawURL
}

// IsValidURL checks if a string is a valid HTTP/HTTPS URL.
func IsValidURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

// GetHealthStatus returns a simple status JSON map for health checks.
func GetHealthStatus() map[string]string {
	return map[string]string{
		"status":  "ok",
		"service": "aeon-mltb-go",
	}
}

// FormatLogMessage formats a log entry for the Go service.
func FormatLogMessage(msg string) string {
	return fmt.Sprintf("[Aeon-Go] %s", msg)
}
