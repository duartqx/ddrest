package main

import (
	"ddrest/models"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("views/index.html"))
		tmpl.Execute(w, nil)
	})

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.Dir("./public"),
			),
		),
	)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		var estate models.Estate

		err := json.NewDecoder(r.Body).Decode(&estate)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Received user data: %+v\n", estate)

		// Marshal the struct to JSON
		jsonData, err := json.Marshal(estate)
		if err != nil {
			http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
			return
		}

		// Set the content type and write the JSON data to the response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	log.Println("Listening and serving on http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
