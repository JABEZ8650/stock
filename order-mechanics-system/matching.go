package main

import (
	"sort"
)

func (ob *OrderBook) AddOrder(o *Order) {
	switch o.Type {
	case Limit:
		if o.Side == Buy {
			ob.Bids = append(ob.Bids, o)
			sort.SliceStable(ob.Bids, func(i, j int) bool {
				if ob.Bids[i].Price == ob.Bids[j].Price {
					return ob.Bids[i].Timestamp.Before(ob.Bids[j].Timestamp)
				}
				return ob.Bids[i].Price > ob.Bids[j].Price
			})
		} else {
			ob.Asks = append(ob.Asks, o)
			sort.SliceStable(ob.Asks, func(i, j int) bool {
				if ob.Asks[i].Price == ob.Asks[j].Price {
					return ob.Asks[i].Timestamp.Before(ob.Asks[j].Timestamp)
				}
				return ob.Asks[i].Price < ob.Asks[j].Price
			})
		}
		ob.MatchOrders()

	case Market:
		ob.ExecuteMarketOrder(o)

	case Stop:
		ob.PendingStops = append(ob.PendingStops, o)
	}
}

func (ob *OrderBook) ExecuteMarketOrder(o *Order) {
	if o.Side == Buy {
		for len(ob.Asks) > 0 && o.Quantity > 0 {
			ask := ob.Asks[0]
			tradeQty := min(o.Quantity, ask.Quantity)
			o.Quantity -= tradeQty
			ask.Quantity -= tradeQty

			if ask.Quantity == 0 {
				ob.Asks = ob.Asks[1:]
			}
		}
	} else {
		for len(ob.Bids) > 0 && o.Quantity > 0 {
			bid := ob.Bids[0]
			tradeQty := min(o.Quantity, bid.Quantity)
			o.Quantity -= tradeQty
			bid.Quantity -= tradeQty

			if bid.Quantity == 0 {
				ob.Bids = ob.Bids[1:]
			}
		}
	}
}

func (ob *OrderBook) MatchOrders() {
	for len(ob.Bids) > 0 && len(ob.Asks) > 0 {
		if ob.Bids[0].Price >= ob.Asks[0].Price {
			bid := ob.Bids[0]
			ask := ob.Asks[0]

			tradeQty := min(bid.Quantity, ask.Quantity)
			bid.Quantity -= tradeQty
			ask.Quantity -= tradeQty

			if bid.Quantity == 0 {
				ob.Bids = ob.Bids[1:]
			}
			if ask.Quantity == 0 {
				ob.Asks = ob.Asks[1:]
			}
		} else {
			break
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
