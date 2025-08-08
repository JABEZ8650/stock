# 📈 Equity Data API

A simple mock API built in Go that simulates real-time equity (stock) data. It provides random equity entries and supports filtering for gainers and losers based on their percentage change.

---

## 🚀 Features

- `GET /api/equities` endpoint
- Mock equity data including:
  - `symbol`: stock symbol (e.g., `ETHB`)
  - `name`: company name (e.g., `EthioBank`)
  - `price`: random price between 100 and 500
  - `changePercent`: daily percentage change (random between -5% and +5%)
  - `volume`: random trade volume
- Query filtering:
  - `?type=gainer` → only returns equities with positive change
  - `?type=loser` → only returns equities with negative change

---

## 📦 Sample Response

```json
[
  {
    "symbol": "ETHB",
    "name": "EthioBank",
    "price": 236.45,
    "changePercent": 1.23,
    "volume": 1345
  },
  {
    "symbol": "ABY",
    "name": "Abyssinia Corp",
    "price": 312.78,
    "changePercent": -2.87,
    "volume": 7480
  }
]
```

##🛠️ How to Run

### 1. Clone the Repository

```
git clone https://github.com/jabez8650/stock.git
cd stock/equity-api
```

### 2. Run the Server

```
go run main.go
```
The server starts on:
📍 http://localhost:8080/api/equities

### 🔎 API Usage

Get All Equities
```
GET /api/equities
```

Filter by Gainers
```
GET /api/equities?type=gainer
```

Filter by Losers
```
GET /api/equities?type=loser
```

## 📁 Project Structure

```
.
├── main.go        # Main application file
├── go.mod         # Go module file
└── README.md      # Project documentation
```