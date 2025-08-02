package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Bond struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Issuer       string  `json:"issuer"`
	CouponRate   float64 `json:"couponRate"`   // in %
	MaturityDate string  `json:"maturityDate"` // e.g., "2029-12-31"
	Price        float64 `json:"price"`        // in USD
}

var bonds = []Bond{
	{"BND001", "Government 5Y", "National Bank", 3.2, "2029-01-01", 102.5},
	{"BND002", "Corporate Alpha", "Ethio Corp", 4.8, "2028-06-15", 98.7},
	{"BND003", "Savings Bond A", "United Trust", 2.5, "2030-12-01", 100.0},
	{"BND004", "Corporate Beta", "Dashen Ltd", 5.0, "2027-09-10", 101.2},
	{"BND005", "Development Bond", "Gov Project", 3.9, "2026-03-20", 97.0},
	{"BND006", "Corporate Gamma", "Abyssinia Co", 6.1, "2031-11-25", 105.3},
	{"BND007", "Municipal Bond", "Addis City", 4.3, "2025-12-31", 99.5},
	{"BND008", "Infrastructure X", "EthDev Fund", 3.5, "2032-07-01", 103.8},
}

func bondSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse query params
	minCoupon, _ := strconv.ParseFloat(r.URL.Query().Get("minCoupon"), 64)
	maxCoupon, _ := strconv.ParseFloat(r.URL.Query().Get("maxCoupon"), 64)
	maturityAfter := r.URL.Query().Get("maturityAfter")
	maturityBefore := r.URL.Query().Get("maturityBefore")

	layout := "2006-01-02"
	var afterDate, beforeDate time.Time
	var err error

	if maturityAfter != "" {
		afterDate, err = time.Parse(layout, maturityAfter)
		if err != nil {
			http.Error(w, "Invalid maturityAfter format (expected YYYY-MM-DD)", http.StatusBadRequest)
			return
		}
	}
	if maturityBefore != "" {
		beforeDate, err = time.Parse(layout, maturityBefore)
		if err != nil {
			http.Error(w, "Invalid maturityBefore format (expected YYYY-MM-DD)", http.StatusBadRequest)
			return
		}
	}

	var filtered []Bond
	for _, bond := range bonds {
		bondDate, err := time.Parse(layout, bond.MaturityDate)
		if err != nil {
			continue // skip malformed
		}

		// Apply filters
		if minCoupon != 0 && bond.CouponRate < minCoupon {
			continue
		}
		if maxCoupon != 0 && bond.CouponRate > maxCoupon {
			continue
		}
		if !afterDate.IsZero() && bondDate.Before(afterDate) {
			continue
		}
		if !beforeDate.IsZero() && bondDate.After(beforeDate) {
			continue
		}

		filtered = append(filtered, bond)
	}

	json.NewEncoder(w).Encode(filtered)
}

func main() {
	http.HandleFunc("/api/bonds/search", bondSearchHandler)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
