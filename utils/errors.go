package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func RaiseMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	err := errors.New(fmt.Sprintf("%s %s Not Allowed", r.URL, r.Method)).Error()
	log.Println(err)
	http.Error(w, err, http.StatusMethodNotAllowed)
}
