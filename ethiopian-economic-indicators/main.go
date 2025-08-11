package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Raw economic data from JSON
type RawEconomicData struct {
	Year             int     `json:"year"`
	DepositRate      float64 `json:"deposit_rate"`
	InflationRate    float64 `json:"inflation_rate"`
	GDPGrowth        float64 `json:"gdp_growth"`
	UnemploymentRate float64 `json:"unemployment_rate"`
	ExchangeRate     float64 `json:"exchange_rate_birr_usd"`
	TradeBalance     float64 `json:"trade_balance_usd"`
}

// Derived indicator structure
type Indicator struct {
	Name        string  `json:"name"`
	Year        int     `json:"year"`
	Value       float64 `json:"value"`
	Unit        string  `json:"unit"`
	Description string  `json:"description"`
}

var rawData []RawEconomicData

func loadEconomicData() {
	file, err := os.Open("economic_data.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&rawData); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}
}

func calculateIndicators(w http.ResponseWriter, r *http.Request) {
	var results []Indicator

	for _, data := range rawData {
		results = append(results, Indicator{
			Name:        "Real Interest Rate",
			Year:        data.Year,
			Value:       data.DepositRate - data.InflationRate,
			Unit:        "%",
			Description: "Deposit rate minus inflation rate; measures the real return on deposits.",
		})

		results = append(results, Indicator{
			Name:        "GDP Growth",
			Year:        data.Year,
			Value:       data.GDPGrowth,
			Unit:        "%",
			Description: "Annual growth rate of Gross Domestic Product.",
		})

		results = append(results, Indicator{
			Name:        "Unemployment Rate",
			Year:        data.Year,
			Value:       data.UnemploymentRate,
			Unit:        "%",
			Description: "Percentage of the labor force that is unemployed.",
		})

		results = append(results, Indicator{
			Name:        "Exchange Rate (ETB/USD)",
			Year:        data.Year,
			Value:       data.ExchangeRate,
			Unit:        "ETB per USD",
			Description: "Average annual exchange rate of Ethiopian Birr per US Dollar.",
		})

		results = append(results, Indicator{
			Name:        "Trade Balance",
			Year:        data.Year,
			Value:       data.TradeBalance,
			Unit:        "Million USD",
			Description: "Net exports (exports minus imports) in million USD.",
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	loadEconomicData()

	r := mux.NewRouter()
	r.HandleFunc("/api/indicators", calculateIndicators).Methods("GET")

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
