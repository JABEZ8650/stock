package main

import (
	"log"
	"net/http"

	myhandlers "github.com/jabez8650/cmsps-api/handlers"

	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	myhandlers.LoadCMSPData()

	// API routes
	r.HandleFunc("/api/cmsps", myhandlers.GetCMSPs).Methods("GET")
	r.HandleFunc("/api/cmsps/{id}", myhandlers.GetCMSPByID).Methods("GET")

	// Enable CORS
	headers := gorillahandlers.AllowedHeaders([]string{"Content-Type"})
	methods := gorillahandlers.AllowedMethods([]string{"GET"})
	origins := gorillahandlers.AllowedOrigins([]string{"*"})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", gorillahandlers.CORS(headers, methods, origins)(r)))
}
