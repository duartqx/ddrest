package main

import (
	"log"
	"net/http"
)

func main() {

	DefineRoutes()

	log.Println("Listening and serving on http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
