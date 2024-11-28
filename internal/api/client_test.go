package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sheldon/internal/config"
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

func TestGetStockDataHistory(t *testing.T) {
	// Mock response data
	mockResponse := `{
		"Time Series (Daily)": {
			"2023-11-26": {
				"1. open": "150.00",
				"2. high": "155.00",
				"3. low": "149.50",
				"4. close": "152.00",
				"5. volume": "1234567"
			},
			"2023-11-25": {
				"1. open": "148.00",
				"2. high": "152.00",
				"3. low": "147.50",
				"4. close": "150.00",
				"5. volume": "987654"
			}
		}
	}`

	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request URL
		expectedURL := "/?function=TIME_SERIES_DAILY&symbol=AAPL&apikey=test_key"
		if r.URL.String() != expectedURL {
			t.Errorf("Expected URL %s, got %s", expectedURL, r.URL.String())
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Create an API client with the mock server
	client := &Client{
		APIKey:     "test_key",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}

	// Test valid case: Fetch daily historical data
	t.Run("FetchDailyHistoricalData", func(t *testing.T) {
		data, err := client.GetStockDataHistory("AAPL", 2, "daily")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Validate the symbol
		if data["symbol"] != "AAPL" {
			t.Errorf("Expected symbol AAPL, got %v", data["symbol"])
		}

		// Validate the series length
		series := data["series"].(map[string]interface{})
		if len(series) != 2 {
			t.Errorf("Expected 2 entries in series, got %d", len(series))
		}

		// Validate a specific entry
		entry := series["2023-11-26"].(map[string]interface{})
		if entry["1. open"] != "150.00" {
			t.Errorf("Expected open value 150.00, got %v", entry["1. open"])
		}
	})

	// Test invalid delta
	t.Run("InvalidDelta", func(t *testing.T) {
		_, err := client.GetStockDataHistory("AAPL", 2, "hourly")
		if err == nil {
			t.Fatal("Expected an error for invalid delta, got nil")
		}
		expectedError := "invalid delta: hourly. Valid options are '1min', '5min', '15min', '30min', '60min', 'daily', 'weekly', 'monthly'"
		if err.Error() != expectedError {
			t.Errorf("Expected error %q, got %q", expectedError, err.Error())
		}
	})

	// Test server returning an unexpected response
	t.Run("UnexpectedResponse", func(t *testing.T) {
		mockInvalidResponse := `{
			"Invalid Key": {}
		}`

		// Update the server to return the invalid response
		server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockInvalidResponse))
		})

		_, err := client.GetStockDataHistory("AAPL", 2, "daily")
		if err == nil {
			t.Fatal("Expected an error for unexpected response format, got nil")
		}
		expectedError := "unexpected response format: map[Invalid Key:map[]]"
		if err.Error() != expectedError {
			t.Errorf("Expected error %q, got %q", expectedError, err.Error())
		}
	})
}
