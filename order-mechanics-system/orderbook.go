package main

type OrderBook struct {
	Bids         []*Order
	Asks         []*Order
	PendingStops []*Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		Bids:         []*Order{},
		Asks:         []*Order{},
		PendingStops: []*Order{},
	}
}
