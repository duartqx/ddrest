package main

import (
	"database/sql"
	"ddrest/models"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"ddrest/controllers"
)

func GetMux() *http.ServeMux {

	indexView := template.Must(template.ParseFiles("views/index.html"))
	estateTableView := template.Must(template.ParseFiles("views/estates-table.html"))

	mux := http.NewServeMux()

	mux.Handle("/", controllers.GetIndexController(indexView))
	mux.Handle("/filterEstate", controllers.GetFilterRealEstateController(estateTableView))

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.Dir("./public"),
			),
		),
	)

	return mux
}

func main() {
	DB, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	models.InitRepository(DB)

	mux := GetMux()

	log.Println("Listening and serving @ http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
