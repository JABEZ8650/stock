package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jabez8650/cmsps-api/models"
)

var cmsps []models.CMSP

func LoadCMSPData() {
	file, _ := os.ReadFile("data/cmsps.json")
	json.Unmarshal(file, &cmsps)
}

func GetCMSPs(w http.ResponseWriter, r *http.Request) {
	typeFilter := r.URL.Query().Get("type")
	licensedBefore := r.URL.Query().Get("licensedBefore")
	licensedAfter := r.URL.Query().Get("licensedAfter")

	filtered := []models.CMSP{}
	for _, c := range cmsps {
		if typeFilter != "" && c.Type != typeFilter {
			continue
		}
		ld, _ := time.Parse("2006-01-02", c.LicensedDate)

		if licensedBefore != "" {
			lb, _ := time.Parse("2006-01-02", licensedBefore)
			if ld.After(lb) {
				continue
			}
		}
		if licensedAfter != "" {
			la, _ := time.Parse("2006-01-02", licensedAfter)
			if ld.Before(la) {
				continue
			}
		}
		filtered = append(filtered, c)
	}
	json.NewEncoder(w).Encode(filtered)
}

func GetCMSPByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, c := range cmsps {
		if c.ID == id {
			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.NotFound(w, r)
}
