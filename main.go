package main

import (
	"log"
	"net/http"

	"tracy_hub/app"
)

func main() {
	http.HandleFunc("/", app.HomeHandler)
	http.HandleFunc("/courses", app.CoursesHandler)
	http.HandleFunc("/sales-enablement", app.SalesEnablementHandler)
	http.HandleFunc("/proposal-templates", app.ProposalTemplatesHandler)
	http.HandleFunc("/pricing", app.PricingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
