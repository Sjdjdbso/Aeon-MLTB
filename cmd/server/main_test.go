package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetListenPort(t *testing.T) {
	os.Unsetenv("GO_PORT")
	os.Unsetenv("PORT")

	if p := getListenPort(); p != "8080" {
		t.Errorf("expected default 8080, got %s", p)
	}

	os.Setenv("PORT", "9090")
	if p := getListenPort(); p != "9090" {
		t.Errorf("expected PORT 9090, got %s", p)
	}

	os.Setenv("GO_PORT", "7070")
	if p := getListenPort(); p != "7070" {
		t.Errorf("expected GO_PORT 7070, got %s", p)
	}

	os.Unsetenv("GO_PORT")
	os.Unsetenv("PORT")
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, status)
	}

	var res map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to parse json response: %v", err)
	}

	if res["status"] != "ok" {
		t.Errorf("expected status ok, got %v", res["status"])
	}
}

func TestStatusHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/status", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(statusHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, status)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to parse json response: %v", err)
	}

	if res["status"] != "running" {
		t.Errorf("expected status running, got %v", res["status"])
	}
}

func TestFilesHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/app/files", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(filesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, status)
	}

	var res map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatalf("failed to parse json response: %v", err)
	}

	if res["engine"] != "golang" {
		t.Errorf("expected engine golang, got %v", res["engine"])
	}
}

func TestSetupRoutes(t *testing.T) {
	mux := setupRoutes()
	ts := httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatalf("failed to make request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", res.StatusCode)
	}
}
