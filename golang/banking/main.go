package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {

	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)

	// starting the sever
	log.Fatal(
		http.ListenAndServe("localhost:8081", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Samuel Puello", City: "Santo domingo", Zipcode: "20101"},
		{Name: "Jose Mejia", City: "Santo domingo", Zipcode: "20102"},
		{Name: "Pedro Sanchez", City: "Santo domingo", Zipcode: "20104"},
	}
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}
