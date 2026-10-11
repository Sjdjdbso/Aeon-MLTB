package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetPort(t *testing.T) {
	origGoPort := os.Getenv("GO_PORT")
	origPort := os.Getenv("PORT")
	defer func() {
		os.Setenv("GO_PORT", origGoPort)
		os.Setenv("PORT", origPort)
	}()

	os.Setenv("GO_PORT", "9090")
	os.Setenv("PORT", "7070")
	if port := getPort(); port != "9090" {
		t.Errorf("expected 9090 when GO_PORT is set, got %s", port)
	}

	os.Unsetenv("GO_PORT")
	os.Setenv("PORT", "7070")
	if port := getPort(); port != "7070" {
		t.Errorf("expected 7070 when PORT is set and GO_PORT is unset, got %s", port)
	}

	os.Unsetenv("PORT")
	if port := getPort(); port != "8080" {
		t.Errorf("expected 8080 when both are unset, got %s", port)
	}
}

func TestHandlers(t *testing.T) {
	mux := setupMux()

	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Root endpoint",
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "Aeon MLTB Go Server is running!",
		},
		{
			name:           "NotFound endpoint",
			path:           "/nonexistent",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found",
		},
		{
			name:           "Health endpoint",
			path:           "/health",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ok","service":"aeon-mltb-go"}`,
		},
		{
			name:           "Files endpoint",
			path:           "/app/files",
			expectedStatus: http.StatusOK,
			expectedBody:   "Aeon MLTB File Manager",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			mux.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			if !strings.Contains(rec.Body.String(), tt.expectedBody) {
				t.Errorf("expected body to contain %q, got %q", tt.expectedBody, rec.Body.String())
			}
		})
	}
}
