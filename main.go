package main

import (
	"log"
	"net/http"

	"github.com/alta3/sales_hub/app"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", app.HomeHandler).Methods("GET")
	router.HandleFunc("/courses", app.CoursesHandler).Methods("GET")
	router.HandleFunc("/sales-enablement", app.SalesEnablementHandler).Methods("GET")
	router.HandleFunc("/proposal-templates", app.ProposalTemplatesHandler).Methods("GET")
	router.HandleFunc("/pricing", app.PricingHandler).Methods("GET")
}

func main() {
	port := "8080" // Change this to your desired port

	r := mux.NewRouter()
	RegisterRoutes(r)

	log.Println("Server is starting on port", port+"...")
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
