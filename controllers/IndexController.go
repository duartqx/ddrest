package controllers

import (
	"ddrest/controllers/utils"
	"html/template"
	"net/http"
)

type IndexController struct{}

func (i IndexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		i.Get(w, r)
	default:
		utils.RaiseMethodNotAllowed(w, r)
	}
}

func (i IndexController) Get(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl.Execute(w, nil)
}
