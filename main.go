package main

import (
	"database/sql"
	"ddrest/models"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var DB *sql.DB

	log.Println("Starting")
	DB, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Println("Error on db connection")
		log.Fatal(err)
	}
	defer DB.Close()

	log.Println("Init Repo")
	models.InitRepository(DB)

	log.Println("Listening and serving @ http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", GetMux()))
}
