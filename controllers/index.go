package controllers

import (
	"html/template"
	"net/http"
)

type IndexController struct{}

func (c IndexController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.Get(w, r)
	case "POST":
		c.Post(w, r)
	default:
		raiseMethodNotAllowed(w, r)
	}
}

func (c *IndexController) Get(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl.Execute(w, nil)
}

func (c *IndexController) Post(w http.ResponseWriter, r *http.Request) {
	raiseMethodNotAllowed(w, r)
}
