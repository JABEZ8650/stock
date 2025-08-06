package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PortfolioRequest struct {
	Profile        string  `json:"profile"`
	InitialCapital float64 `json:"initialCapital"`
	TargetGoal     float64 `json:"targetGoal"`
}

type Asset struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Yield  float64 `json:"yield"`
}

type PortfolioResponse struct {
	Stock       Asset   `json:"stock"`
	Bond        Asset   `json:"bond"`
	Cash        Asset   `json:"cash"`
	TotalReturn float64 `json:"totalReturn"`
	GoalMet     bool    `json:"goalMet"`
	GapToGoal   float64 `json:"gapToGoal"`
}

func recommendPortfolio(c *gin.Context) {
	var req PortfolioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var stockPct, bondPct, cashPct float64

	switch req.Profile {
	case "active":
		stockPct = 0.6
		bondPct = 0.3
		cashPct = 0.1
	case "passive":
		stockPct = 0.3
		bondPct = 0.4
		cashPct = 0.3
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile"})
		return
	}

	capital := req.InitialCapital

	stockAmt := capital * stockPct
	bondAmt := capital * bondPct
	cashAmt := capital * cashPct

	stockYield := stockAmt * 0.20
	bondYield := bondAmt * 0.10
	cashYield := cashAmt * 0.07

	totalReturn := stockYield + bondYield + cashYield
	gap := totalReturn - req.TargetGoal
	goalMet := totalReturn >= req.TargetGoal

	response := PortfolioResponse{
		Stock:       Asset{"EthioTelecom", stockAmt, stockYield},
		Bond:        Asset{"Ethiopian Gov Bond", bondAmt, bondYield},
		Cash:        Asset{"Birr", cashAmt, cashYield},
		TotalReturn: totalReturn,
		GoalMet:     goalMet,
		GapToGoal:   gap,
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()
	r.POST("/api/portfolio/recommend", recommendPortfolio)
	r.Run(":8080")
}
