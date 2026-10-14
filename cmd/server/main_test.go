package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	os.Unsetenv("GO_PORT")
	os.Unsetenv("PORT")

	if p := getPort(); p != "8080" {
		t.Errorf("expected default port 8080, got %s", p)
	}

	os.Setenv("PORT", "3000")
	if p := getPort(); p != "3000" {
		t.Errorf("expected PORT 3000, got %s", p)
	}

	os.Setenv("GO_PORT", "9090")
	if p := getPort(); p != "9090" {
		t.Errorf("expected GO_PORT 9090, got %s", p)
	}

	os.Unsetenv("GO_PORT")
	os.Unsetenv("PORT")
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	healthHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != "OK" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), "OK")
	}
}

func TestStatusHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/status", nil)
	rr := httptest.NewRecorder()

	statusHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp StatusResponse
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "running" {
		t.Errorf("expected status 'running', got '%s'", resp.Status)
	}
}
