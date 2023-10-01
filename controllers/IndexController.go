package controllers

import (
	"ddrest/utils"
	"html/template"
	"net/http"
)

type IndexController struct{}

func (i IndexController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		i.Get(w, r)
	case "POST":
		i.Post(w, r)
	default:
		utils.RaiseMethodNotAllowed(w, r)
	}
}

func (i IndexController) Get(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl.Execute(w, nil)
}

func (i IndexController) Post(w http.ResponseWriter, r *http.Request) {
	utils.RaiseMethodNotAllowed(w, r)
}