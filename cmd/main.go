package main

import (
	"sheldon/internal/api"
	"sheldon/internal/config"
	"sheldon/internal/handler"
	"flag"
	"fmt"
	"log"
)

func main() {
	// Define command-line flags
	symbol := flag.String("symbol", "", "Stock symbol to query (e.g., IBM, NVDA)")
	flag.Parse()

	// Validate input
	if *symbol == "" {
		log.Fatal("Errror: stock symbol is requried. user the -symbol flag to specify it.")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Create Api Client and query handler
	client := api.NewClient(cfg)
	queryHandler := handler.NewQueryHandler(client)
	// Fetch stock data
	jsonData, err := queryHandler.HandleStockQuery(*symbol)
	if err != nil {
		log.Fatalf("Error fetching stock data: %v", err)
	}

	// Output the result
	fmt.Println(jsonData)

}
