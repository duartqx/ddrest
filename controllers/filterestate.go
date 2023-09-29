package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type FilterEstateController struct {
	Name          string `json:"estate_name"`
	AddressNumber string `json:"estate_address_number"`
	Floors        int    `json:"estate_floors"`
	Balconies     int    `json:"estate_balconies"`
	Bedrooms      int    `json:"estate_bedrooms"`
	Bathrooms     int    `json:"estate_bathrooms"`
	PetsAllowed   bool   `json:"estate_pets_allowed"`
}

func (f FilterEstateController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		f.Get(w, r)
	case "POST":
		f.Post(w, r)
	default:
		raiseMethodNotAllowed(w, r)
	}
}

func (f FilterEstateController) Get(w http.ResponseWriter, r *http.Request) {
	raiseMethodNotAllowed(w, r)
}

func (f FilterEstateController) Post(w http.ResponseWriter, r *http.Request) {

	var filter FilterEstateController

	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		log.Println(r.URL, r.Method, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%s %s Received user data: %+v\n", r.URL, r.Method, filter)

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(filter)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	// Set the content type and write the JSON data to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
