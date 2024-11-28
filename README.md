
# **Sheldon CLI**

SHELDON: 'Stock Handling and Evaluation Library for Data Operations and Navigation'

Concept:

**SHELDON** is a command-line interface (CLI) tool aimed at helping users analyze, manage, and navigate financial data. It could support tasks like:
- Fetching and analyzing stock prices.
- Evaluating portfolio performance.
- Navigating datasets or APIs related to financial markets.
- Performing computations like ROI, CAGR, or risk assessment.


---

## **Features**
- Query stock prices for any publicly traded company by its symbol.
- Output data in JSON format for easy integration or further analysis.
- Modular design for future extensibility (e.g., scheduled queries, database integrations).

---

## **Requirements**
- [Go](https://golang.org/doc/install) (version 1.17 or later)
- An [Alpha Vantage API key](https://www.alphavantage.co/support/#api-key)
- Internet connection

---

## **Installation**

1. Clone the repository:
   ```bash
   git clone git@github.com:Cesarec88/sheldon.git
   cd sheldon
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up your `.env` file:
   Create a `.env` file in the project root and add your API key and the base URL:
   ```
   ALPHA_VANTAGE_API_KEY=your_api_key_here
   BASE_URL=https://www.alphavantage.co/query
   ```

4. Build the application:
   ```bash
   go build -o sheldon cmd/main.go
   ```

---

## **Usage**

Run the CLI with the `-symbol` flag to query stock data:
```bash
./sheldon -symbol=<stock_symbol>
```

### **Example**
```bash
./sheldon -symbol=IBM
```
Output:
```json
{
  "Global Quote": {
    "01. symbol": "IBM",
    "05. price": "123.45"
  }
}
```

---

## **Testing**

Run the unit tests:
```bash
go test ./...
```

Expected output:
```
ok  	sheldon/internal/config   0.XXXs
ok  	sheldon/internal/api      0.XXXs
ok  	sheldon/internal/handler  0.XXXs
```

---

## **Project Structure**
```
sheldon/
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
```

---

## **Future Enhancements**
- **Scheduled Queries**: Add a feature to fetch stock data periodically.
- **Multiple Stock Support**: Query multiple stocks at once.
- **Output Options**: Save results to a file (JSON, CSV) or database.
- **Web Interface**: Expand to a web application for broader accessibility.

---

## **License**
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## **Contributing**
Contributions are welcome! Feel free to open issues or submit pull requests to improve this project.

---

## **Acknowledgments**
- [Alpha Vantage](https://www.alphavantage.co/) for providing the stock market data API.
- [Go](https://golang.org/) for powering this application.
