package main

import (
	"ddrest/controllers"
	"net/http"
)

func GetMux() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.IndexController{}.Dispatch)
	mux.HandleFunc("/filterEstate", controllers.FilterRealEstateController{}.Dispatch)

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
