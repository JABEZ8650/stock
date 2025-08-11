# Ethiopian Economic Indicator API

A simple Go REST API that serves Ethiopian macroeconomic data (2019–2024) and calculates derived indicators such as the Real Interest Rate.

## Features
- Loads static macroeconomic data from `economic_data.json`
- Calculates **Real Interest Rate** = Deposit Rate − Inflation
- Includes 4 additional indicators:
  - GDP Growth
  - Unemployment Rate
  - Exchange Rate (ETB/USD)
  - Trade Balance (USD)
- Year-by-year data (2019–2024)
- Returns metadata for each indicator (name, year, value, unit, description)
- No database or external API needed

## Project Structure

ethiopian-economic-indicators/
├── economic_data.json # Raw macroeconomic data
├── main.go # API implementation
├── go.mod # Go module definition
└── README.md # Project documentation


## Installation & Running
1. Clone the repository:
   ```bash
   git clone https://github.com/<your-github-username>/ethiopian-economic-indicators.git
   cd ethiopian-economic-indicators
   ```

   ```
   go mod tidy
   ```

   ```
   go run main.go
   ```

   ```
   GET http://localhost:8080/api/indicators
   ```

### peace out!!!