package cmd

import (
	"fmt"
	"log"
	"sheldon/internal/api"
	"sheldon/internal/config"
	"sheldon/internal/handler"

	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch stock data",
	Long: `Fetch the latest stock quote or historical data for a given ticker symbol.
Example usage:
  sheldon fetch --ticker AAPL`,
	Run: func(cmd *cobra.Command, args []string) {
		// Parse flags
		ticker, _ := cmd.Flags().GetString("ticker")
		limit, _ := cmd.Flags().GetInt("limit")
		delta, _ := cmd.Flags().GetString("delta")

		// Validate required flags
		if ticker == "" {
			log.Fatal("Error: --ticker flag is required")
		}

		// Load configuration
		cfg := config.LoadConfig()
		client := api.NewClient(cfg)
		queryHandler := handler.NewQueryHandler(client)

		var jsonData string
		var err error

		// Fetch data
		if limit == 1 && delta == "" {
			// Fetch the latest stock quote
			jsonData, err = queryHandler.GetStockQuote(ticker)
		} else {
			// Fetch historical data
			jsonData, err = queryHandler.GetStockDataHistory(ticker, limit, delta)
		}

		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}

		// Output the result
		fmt.Println(jsonData)
	},
}

func init() {
	// Add fetch command to the root command
	rootCmd.AddCommand(fetchCmd)

	// Define flags for the fetch command
	fetchCmd.Flags().String("ticker", "", "Stock ticker to query (e.g., AAPL)")
	fetchCmd.Flags().Int("limit", 1, "Number of historical elements to return (default: 1 for latest data)")
	fetchCmd.Flags().String("delta", "", "Time delta for historical data (e.g., '1min', 'daily', 'weekly', 'monthly')")
}
