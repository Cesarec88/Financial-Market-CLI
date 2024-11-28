package handler

import (
	"encoding/json"
	"fmt"
	"sheldon/internal/api"
)

type QueryHandler struct {
	APIClient *api.Client
}

func NewQueryHandler(client *api.Client) *QueryHandler {
	return &QueryHandler{APIClient: client}
}

func (qh QueryHandler) GetStockQuote(symbol string) (string, error) {
	// Fetch stock data
	data, err := qh.APIClient.GetStockQuote(symbol)
	if err != nil {
		return "", fmt.Errorf("error fetching stock data: %v", err)
	}

	// Format the response as a pretty JSON
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", fmt.Errorf("error formatting stock data: %v", err)
	}

	return string(jsonData), nil
}

// GetStockDataHistory fetches historical stock data for a given ticker
func (qh *QueryHandler) GetStockDataHistory(symbol string, limit int, delta string) (string, error) {
	data, err := qh.APIClient.GetStockDataHistory(symbol, limit, delta)
	if err != nil {
		return "", err
	}

	// Format the response as pretty JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
