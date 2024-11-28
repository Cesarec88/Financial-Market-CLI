package handler

import (
	"encoding/json"
	"sheldon/internal/api"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandleStockQuery(t *testing.T) {
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
	jsonData, err := queryHandler.HandleStockQuery("IBM")
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
