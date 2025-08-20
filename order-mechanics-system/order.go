package main

import "time"

type OrderSide string
type OrderType string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"

	Market OrderType = "market"
	Limit  OrderType = "limit"
	Stop   OrderType = "stop"
)

type Order struct {
	ID        string    `json:"id"`
	Side      OrderSide `json:"side"`
	Type      OrderType `json:"type"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}
