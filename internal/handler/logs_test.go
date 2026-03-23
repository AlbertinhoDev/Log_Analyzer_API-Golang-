package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"loganalyzerapi/internal/service"
)

func TestLogsHandlerHealth(t *testing.T) {
	parser := service.NewParser()
	analyzer := service.NewAnalyzer(parser)
	handler := NewLogsHandler(analyzer)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `"status":"ok"`) {
		t.Fatalf("expected response body to contain status ok, got %s", body)
	}
}

func TestLogsHandlerAnalyze(t *testing.T) {
	parser := service.NewParser()
	analyzer := service.NewAnalyzer(parser)
	handler := NewLogsHandler(analyzer)

	requestBody := `{
		"log_text": "2026-03-20 10:15:22 ERROR Connection timeout\n2026-03-20 10:15:25 INFO Retry request"
	}`

	req := httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Analyze(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()

	if !strings.Contains(body, `"total_lines":2`) {
		t.Fatalf("expected response body to contain total_lines=2, got %s", body)
	}

	if !strings.Contains(body, `"error_count":1`) {
		t.Fatalf("expected response body to contain error_count=1, got %s", body)
	}

	if !strings.Contains(body, `"info_count":1`) {
		t.Fatalf("expected response body to contain info_count=1, got %s", body)
	}

	if !strings.Contains(body, `"top_messages"`) {
		t.Fatalf("expected response body to contain top_messages, got %s", body)
	}
}

func TestLogsHandlerAnalyzeInvalidJSON(t *testing.T) {
	parser := service.NewParser()
	analyzer := service.NewAnalyzer(parser)
	handler := NewLogsHandler(analyzer)

	req := httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBufferString("not json"))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler.Analyze(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, `"error":"invalid JSON body"`) {
		t.Fatalf("expected error response, got %s", body)
	}
}
