package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/orders", createOrderHandler)
	http.HandleFunc("/orderbook", getOrderBookHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
