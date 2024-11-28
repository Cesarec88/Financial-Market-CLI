package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sheldon/internal/api"
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Mock API client
	client := &api.Client{
		APIKey:     "test_key",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}

	// Create query handler
	queryHandler := NewQueryHandler(client)

	// Test querying a stock
	jsonData, err := queryHandler.GetStockQuote("IBM")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	// Validate the returned JSON
	expectedJSON := `{
        "Global Quote": {
            "01. symbol": "IBM",
            "05. price": "123.45"
        }
    }`

	// Unmarshal the expected JSON
	var expectedData map[string]interface{}
	err = json.Unmarshal([]byte(expectedJSON), &expectedData)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the actual JSON output
	var actualData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &actualData)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actualData, expectedData) {
		t.Fatalf("JSON data does not match:\nExpected: %v\nGot: %v", expectedData, actualData)
	}

}

func TestQueryHandler_GetStockDataHistory(t *testing.T) {
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	// Create an API client with the mock server
	client := &api.Client{
		APIKey:     "test_key",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}

	// Create a QueryHandler instance
	queryHandler := NewQueryHandler(client)

	// Test valid case: Fetch daily historical data
	t.Run("FetchDailyHistoricalData", func(t *testing.T) {
		jsonData, err := queryHandler.GetStockDataHistory("AAPL", 2, "daily")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Parse the returned JSON
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
			t.Fatalf("Error unmarshaling JSON data: %v", err)
		}

		// Validate the symbol
		if result["symbol"] != "AAPL" {
			t.Errorf("Expected symbol AAPL, got %v", result["symbol"])
		}

		// Validate the series length
		series := result["series"].(map[string]interface{})
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
		_, err := queryHandler.GetStockDataHistory("AAPL", 2, "hourly")
		if err == nil {
			t.Fatal("Expected an error for invalid delta, got nil")
		}
		expectedError := "invalid delta: hourly. Valid options are '1min', '5min', '15min', '30min', '60min', 'daily', 'weekly', 'monthly'"
		if err.Error() != expectedError {
			t.Errorf("Expected error %q, got %q", expectedError, err.Error())
		}
	})
}
