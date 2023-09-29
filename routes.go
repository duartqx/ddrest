package main

import (
	"ddrest/controllers"
	"net/http"
)

func DefineRoutes() {

	http.HandleFunc("/", controllers.IndexController{}.Dispatch)
	http.HandleFunc("/form", controllers.FormController{}.Dispatch)

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.Dir("./public"),
			),
		),
	)
}
