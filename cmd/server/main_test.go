package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetPort(t *testing.T) {
	os.Unsetenv("GO_PORT")
	os.Unsetenv("PORT")

	if p := getPort(); p != "8080" {
		t.Errorf("expected default 8080, got %s", p)
	}

	os.Setenv("PORT", "9090")
	if p := getPort(); p != "9090" {
		t.Errorf("expected PORT 9090, got %s", p)
	}

	os.Setenv("GO_PORT", "7070")
	if p := getPort(); p != "7070" {
		t.Errorf("expected GO_PORT 7070, got %s", p)
	}

	os.Unsetenv("GO_PORT")
	os.Unsetenv("PORT")
}

func TestGetPythonBackendURL(t *testing.T) {
	os.Unsetenv("PYTHON_BACKEND_URL")
	if url := getPythonBackendURL(); url != "http://127.0.0.1:8000" {
		t.Errorf("expected default http://127.0.0.1:8000, got %s", url)
	}

	os.Setenv("PYTHON_BACKEND_URL", "http://localhost:5000")
	if url := getPythonBackendURL(); url != "http://localhost:5000" {
		t.Errorf("expected custom URL http://localhost:5000, got %s", url)
	}
	os.Unsetenv("PYTHON_BACKEND_URL")
}

func TestHealthEndpoint(t *testing.T) {
	mux := setupRoutes()
	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `"status":"ok"`) {
		t.Errorf("unexpected body: %s", body)
	}
}

func TestPingEndpoint(t *testing.T) {
	mux := setupRoutes()
	req := httptest.NewRequest("GET", "/ping", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `"status":"ok"`) {
		t.Errorf("unexpected body: %s", body)
	}
}

func TestHomepageEndpoint(t *testing.T) {
	mux := setupRoutes()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "mirror-leech-telegram-bot") {
		t.Errorf("expected homepage content, got: %s", body)
	}
}

func TestAppFilesEndpoint(t *testing.T) {
	mux := setupRoutes()
	req := httptest.NewRequest("GET", "/app/files", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}
