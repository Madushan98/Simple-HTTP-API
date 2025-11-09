package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"simple-http-api/api/handler"
)

func TestHelloWorldHandler(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		expectedCode   int
		expectedField  string
		expectedValue  string
	}{
		{"ValidNameA", "?name=Alice", http.StatusOK, "message", "Hello Alice"},
		{"ValidNameM", "?name=Michael", http.StatusOK, "message", "Hello Michael"},
		{"InvalidNameZ", "?name=Zane", http.StatusBadRequest, "error", "Invalid Input"},
		{"MissingName", "", http.StatusBadRequest, "error", "Invalid Input"},
		{"EmptyName", "?name=", http.StatusBadRequest, "error", "Invalid Input"},
		{"NonAlphabetic", "?name=123", http.StatusBadRequest, "error", "Invalid Input"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/hello-world"+tc.query, nil)
			rec := httptest.NewRecorder()

			handler.HelloWorldHandler(rec, req)
			res := rec.Result()

			if res.StatusCode != tc.expectedCode {
				t.Errorf("expected %d, got %d", tc.expectedCode, res.StatusCode)
			}

			var body map[string]string
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if val, ok := body[tc.expectedField]; !ok || val != tc.expectedValue {
				t.Errorf("expected %s=%s, got %v", tc.expectedField, tc.expectedValue, body)
			}
		})
	}
}
