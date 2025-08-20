# Order Mechanics System

A simple in-memory order matching engine written in Go using `net/http`.

## Features
- Place Buy/Sell orders (`limit`, `market`, `stop`)
- Match orders automatically
- In-memory order book with bids/asks
- HTTP API interface

## API Endpoints
### `POST /orders`
Place a new order.

### `GET /orderbook`
Returns the current order book.

## Run Locally
```
go run .
```

## Example Usage
```
curl -X POST http://localhost:8080/orders -d '{"id":"1","side":"buy","type":"limit","quantity":100,"price":50}' -H "Content-Type: application/json"
curl http://localhost:8080/orderbook
```

## Peace out!!!