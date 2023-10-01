package controllers

import (
	"ddrest/controllers/utils"
	"ddrest/models"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type RealEstateController struct{}

func (re RealEstateController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		re.Get(w, r)
	case "POST":
		re.Post(w, r)
	default:
		utils.RaiseMethodNotAllowed(w, r)
	}
}

func (re RealEstateController) Post(w http.ResponseWriter, r *http.Request) {

	fileBytes, err := os.ReadFile("MOCK_DATA.json")
	if err != nil {
		log.Fatal(err)
	}

	var realEstates *[]models.RealEstate
	err = json.Unmarshal(fileBytes, &realEstates)
	if err != nil {
		log.Fatal(err)
	}
	models.RealEstate{}.Create(realEstates)
}

func (re RealEstateController) Get(w http.ResponseWriter, r *http.Request) {
	utils.RaiseMethodNotAllowed(w, r)
}
