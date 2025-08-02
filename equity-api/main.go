package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Equity struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	ChangePercent float64 `json:"changePercent"`
	Volume        int     `json:"volume"`
}

func generateMockEquities() []Equity {
	companies := []Equity{
		{"ETHB", "EthioBank", 0, 0, 0},
		{"ABY", "Abyssinia Corp", 0, 0, 0},
		{"ZGBC", "Zemen Group", 0, 0, 0},
		{"DASH", "Dashen Limited", 0, 0, 0},
		{"UBNK", "United Bank", 0, 0, 0},
	}

	rand.Seed(time.Now().UnixNano())
	for i := range companies {
		companies[i].Price = randFloat(100, 500)
		companies[i].ChangePercent = randFloat(-5, 5)
		companies[i].Volume = rand.Intn(10000) + 100
	}

	return companies
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func equitiesHandler(w http.ResponseWriter, r *http.Request) {
	equities := generateMockEquities()

	queryType := strings.ToLower(r.URL.Query().Get("type"))
	if queryType == "gainer" {
		var gainers []Equity
		for _, e := range equities {
			if e.ChangePercent > 0 {
				gainers = append(gainers, e)
			}
		}
		equities = gainers
	} else if queryType == "loser" {
		var losers []Equity
		for _, e := range equities {
			if e.ChangePercent < 0 {
				losers = append(losers, e)
			}
		}
		equities = losers
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equities)
}

func main() {
	http.HandleFunc("/api/equities", equitiesHandler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
