package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShortenHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/shorten", nil)
	rec := httptest.NewRecorder()

	shortenHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", rec.Code)
	}
}

func TestShortenHandler_MissingURL(t *testing.T) {
	body, _ := json.Marshal(map[string]string{"url": ""})
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	shortenHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", rec.Code)
	}
}

func TestGenerateShortCode(t *testing.T) {
	code := generateShortCode(6)

	if len(code) != 6 {
		t.Errorf("Expected code length 6, got %d", len(code))
	}
}

func TestGenerateShortCodeUniqueness(t *testing.T) {
	code1 := generateShortCode(6)
	code2 := generateShortCode(6)

	if code1 == code2 {
		t.Errorf("Expected different codes, got the same: %s", code1)
	}
}
