package api

import (
	"sheldon/internal/config"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetStockQuote(t *testing.T) {
	// Mock API response
	mockResponse := `{
		"Global Quote": {
			"01. symbol": "IBM",
			"05. price": "123.45"
		}
	}`

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request parameters (optional)
		if r.URL.Query().Get("function") != "GLOBAL_QUOTE" {
			t.Fatalf("Expected function GLOBAL_QUOTE, got %s", r.URL.Query().Get("function"))
		}
		if r.URL.Query().Get("symbol") != "IBM" {
			t.Fatalf("Expected symbol IBM, got %s", r.URL.Query().Get("symbol"))
		}

		// Respond with the mock data
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, mockResponse)
	}))
	defer server.Close()

	// Create a mock config that uses the server's URL
	cfg := &config.Config{
		APIKey:  "test_key",
		BaseURL: server.URL, // Use the mock server's URL
	}

	// Create a client using the mock config
	client := &Client{
		APIKey:     cfg.APIKey,
		BaseURL:    cfg.BaseURL,
		HTTPClient: server.Client(),
	}

	// Make the API call
	data, err := client.GetStockQuote("IBM")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate the response
	globalQuote, ok := data["Global Quote"].(map[string]interface{})
	if !ok {
		t.Fatalf("Unexpected response format: %v", data)
	}

	if globalQuote["01. symbol"] != "IBM" {
		t.Fatalf("Expected symbol IBM, got %s", globalQuote["01. symbol"])
	}

	if globalQuote["05. price"] != "123.45" {
		t.Fatalf("Expected price 123.45, got %s", globalQuote["05. price"])
	}
}
