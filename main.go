package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening and serving on http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", GetMux()))
}
