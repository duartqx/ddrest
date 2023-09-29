package controllers

import (
	"ddrest/models"
	"encoding/json"
	"log"
	"net/http"
)

type FormController struct{}

func (c FormController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.Get(w, r)
	case "POST":
		c.Post(w, r)
	default:
		raiseMethodNotAllowed(w, r)
	}
}

func (c *FormController) Get(w http.ResponseWriter, r *http.Request) {
	raiseMethodNotAllowed(w, r)
}

func (c *FormController) Post(w http.ResponseWriter, r *http.Request) {
	var estate models.Estate

	err := json.NewDecoder(r.Body).Decode(&estate)
	if err != nil {
		log.Println(r.URL, r.Method, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%s %s Received user data: %+v\n", r.URL, r.Method, estate)

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(estate)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	// Set the content type and write the JSON data to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
