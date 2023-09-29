package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func raiseMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	err := errors.New(fmt.Sprintf("%s %s Not Allowed", r.URL, r.Method)).Error()
	log.Println(err)
	http.Error(w, err, http.StatusMethodNotAllowed)
}
