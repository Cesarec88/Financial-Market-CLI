package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sheldon/internal/config"
)

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		APIKey:     cfg.APIKey,
		BaseURL:    cfg.BaseURL,
		HTTPClient: &http.Client{},
	}
}

// GetStockQuote fetches the stock data using the given ticker.
func (c *Client) GetStockQuote(symbol string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", c.BaseURL, symbol, c.APIKey)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding API response: %v", err)
	}

	return result, nil
}

// GetStockDataHistory fetches historical stock data using the time-series endpoints.
func (c *Client) GetStockDataHistory(symbol string, limit int, delta string) (map[string]interface{}, error) {
	// Determine the function and interval based on delta
	var function, interval string
	switch delta {
	case "1min", "5min", "15min", "30min", "60min":
		function = "TIME_SERIES_INTRADAY"
		interval = delta
	case "daily":
		function = "TIME_SERIES_DAILY"
	case "weekly":
		function = "TIME_SERIES_WEEKLY"
	case "monthly":
		function = "TIME_SERIES_MONTHLY"
	default:
		return nil, fmt.Errorf("invalid delta: %s. Valid options are '1min', '5min', '15min', '30min', '60min', 'daily', 'weekly', 'monthly'", delta)
	}

	// Build the API URL
	url := fmt.Sprintf("%s?function=%s&symbol=%s&apikey=%s", c.BaseURL, function, symbol, c.APIKey)
	if interval != "" {
		url = fmt.Sprintf("%s&interval=%s", url, interval)
	}

	// Make the API request
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding API response: %v", err)
	}

	// Extract the time series data
	timeSeriesKey := "Time Series (1min)"
	if delta == "daily" {
		timeSeriesKey = "Time Series (Daily)"
	} else if delta == "weekly" {
		timeSeriesKey = "Weekly Time Series"
	} else if delta == "monthly" {
		timeSeriesKey = "Monthly Time Series"
	}

	timeSeries, ok := result[timeSeriesKey].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format: %v", result)
	}

	// Limit the results
	limitedSeries := make(map[string]interface{})
	count := 0
	for date, data := range timeSeries {
		if count >= limit {
			break
		}
		limitedSeries[date] = data
		count++
	}

	return map[string]interface{}{
		"symbol": symbol,
		"series": limitedSeries,
	}, nil
}
