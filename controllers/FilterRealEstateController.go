package controllers

import (
	"ddrest/controllers/utils"
	"ddrest/models"
	"html/template"
	"log"
	"net/http"
)

type FilterRealEstateController struct {
	estateTableView *template.Template
}

func GetFilterRealEstateController(estateTableView *template.Template) *FilterRealEstateController {
	return &FilterRealEstateController{
		estateTableView: estateTableView,
	}
}

func (f FilterRealEstateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		f.Post(w, r)
	default:
		utils.RaiseMethodNotAllowed(w, r)
	}
}

func (f FilterRealEstateController) Post(w http.ResponseWriter, r *http.Request) {

	var re models.RealEstate

	if err := re.Populate(r.Body); err != nil {
		log.Println(r.URL, r.Method, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f.estateTableView.Execute(w, re.Filter())
}
