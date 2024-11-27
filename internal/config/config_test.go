package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a mock .env file
	envContent := `ALPHA_VANTAGE_API_KEY=test_key
					BASE_URL=http://testbroker.com`
	os.WriteFile(".env", []byte(envContent), 0644)
	defer os.Remove(".env")
	cfg := LoadConfig()
	if cfg.APIKey != "test_key" {
		t.Fatalf("Expected API key 'test_key', got '%s'", cfg.APIKey)
	}
	if cfg.BaseURL != "http://testbroker.com" {
		t.Fatalf("Expected BaseURL 'http://testbroker.com', got '%s'", cfg.BaseURL)
	}

}
