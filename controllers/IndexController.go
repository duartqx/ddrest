package controllers

import (
	"ddrest/controllers/utils"
	"html/template"
	"net/http"
)

type IndexController struct {
	indexView *template.Template
}

func GetIndexController(indexView *template.Template) *IndexController {
	return &IndexController{
		indexView: indexView,
	}
}

func (i IndexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		i.Get(w, r)
	default:
		utils.RaiseMethodNotAllowed(w, r)
	}
}

func (i IndexController) Get(w http.ResponseWriter, r *http.Request) {
	i.indexView.Execute(w, nil)
}
