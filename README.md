README for financial-market-cli

Financial Market CLI

The Financial Market CLI is a lightweight, command-line application that fetches stock market data using the Alpha Vantage API. It provides real-time or near real-time stock information in JSON format.

Features

	•	Query stock prices for any publicly traded company by its symbol.
	•	Output data in JSON format for easy integration or further analysis.
	•	Modular design for future extensibility (e.g., scheduled queries, database integrations).

Requirements

	•	Go (version 1.17 or later)
	•	An Alpha Vantage API key
	•	Internet connection

Installation

	1.	Clone the repository:

git clone https://github.com/yourusername/financial-market-cli.git
cd financial-market-cli


	2.	Install dependencies:

go mod tidy


	3.	Set up your .env file:
Create a .env file in the project root and add your API key and the base URL:

ALPHA_VANTAGE_API_KEY=your_api_key_here
BASE_URL=https://www.alphavantage.co/query


	4.	Build the application:

go build -o financial-market-cli cmd/main.go

Usage

Run the CLI with the -symbol flag to query stock data:

./financial-market-cli -symbol=<stock_symbol>

Example

./financial-market-cli -symbol=IBM

Output:

{
  "Global Quote": {
    "01. symbol": "IBM",
    "05. price": "123.45"
  }
}

Testing

Run the unit tests:

go test ./...

Expected output:

ok  	financial-market-cli/internal/config   0.XXXs
ok  	financial-market-cli/internal/api      0.XXXs
ok  	financial-market-cli/internal/handler  0.XXXs

Project Structure

financial-market-cli/
│
├── cmd/                  # Command-line interface
│   └── main.go           # Entry point for the CLI
│
├── internal/             # Internal packages
│   ├── api/              # API client implementations
│   │   └── client.go     # Core API client logic
│   │
│   ├── config/           # Configuration management
│   │   └── config.go     # Loads environment variables and settings
│   │
│   └── handler/          # Query handling logic
│       └── query_handler.go # Processes stock queries
│
├── .env                  # Environment file for sensitive settings
├── go.mod                # Go module definition
├── go.sum                # Dependency tracking
└── README.md             # Project documentation

Future Enhancements

	•	Scheduled Queries: Add a feature to fetch stock data periodically.
	•	Multiple Stock Support: Query multiple stocks at once.
	•	Output Options: Save results to a file (JSON, CSV) or database.
	•	Web Interface: Expand to a web application for broader accessibility.

License

This project is licensed under the MIT License. See the LICENSE file for details.

Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve this project.

Acknowledgments

	•	Alpha Vantage for providing the stock market data API.
	•	Go for powering this application.
