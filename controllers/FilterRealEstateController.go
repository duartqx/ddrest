package controllers

import (
	"ddrest/filters"
	"ddrest/utils"
	"html/template"
	"log"
	"net/http"
)

type FilterRealEstateController struct{}

func (f FilterRealEstateController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		f.Get(w, r)
	case "POST":
		f.Post(w, r)
	default:
		utils.RaiseMethodNotAllowed(w, r)
	}
}

func (f FilterRealEstateController) Post(w http.ResponseWriter, r *http.Request) {

	var filter filters.FilterRealEstateForm

	if err := filter.AddBody(r.Body); err != nil {
		log.Println(r.URL, r.Method, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("views/estates-table.html"))
	tmpl.Execute(w, filter.Filter())
}

func (f FilterRealEstateController) Get(w http.ResponseWriter, r *http.Request) {
	utils.RaiseMethodNotAllowed(w, r)
}