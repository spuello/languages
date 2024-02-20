package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
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
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
