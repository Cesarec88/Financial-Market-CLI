package handler

import (
	"encoding/json"
	"financial-market-cli/internal/api"
	"fmt"
)

type QueryHandler struct {
	APIClient *api.Client
}

func NewQueryHandler(client *api.Client) *QueryHandler {
	return &QueryHandler{APIClient: client}
}

func (qh QueryHandler) HandleStockQuery(symbol string) (string, error) {
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
