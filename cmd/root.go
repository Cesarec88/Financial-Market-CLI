package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sheldon",
	Short: "Sheldon CLI: Analyze, manage, and navigate financial data.",
	Long: `SHELDON: Stock Handling and Evaluation Library for Data Operations and Navigation.

The Sheldon CLI is a command-line tool for analyzing stock data, querying historical records, and navigating financial datasets.
For more details, run 'sheldon help'.`,
	// Optional Run function for the root command:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Sheldon CLI!")
		fmt.Println("Use 'sheldon help' to explore available commands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Local flags are specific to this command.
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
}
