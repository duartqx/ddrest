package controllers

import (
	"ddrest/models"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type FilterForm struct {
	Name          string `json:"estate_name"`
	AddressNumber string `json:"estate_address_number"`
	Floors        int    `json:"estate_floors"`
	Balconies     int    `json:"estate_balconies"`
	Bedrooms      int    `json:"estate_bedrooms"`
	Bathrooms     int    `json:"estate_bathrooms"`
	PetsAllowed   int    `json:"estate_pets_allowed"`
}

func (f *FilterForm) addBody(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(&f)
}

func (f FilterForm) filter(estates *[]models.Estate) (result *[]models.Estate) {
	result = &[]models.Estate{}

	for _, e := range *estates {
		switch {
		case f.Name != "" && e.Name != f.Name:
			continue
		case f.AddressNumber != "" && e.AddressNumber != f.AddressNumber:
			continue
		case f.Floors > e.Floors:
			continue
		case f.Balconies > e.Balconies:
			continue
		case f.Bedrooms > e.Bedrooms:
			continue
		case f.Bathrooms > e.Bathrooms:
			continue
		case f.PetsAllowed != 0 && f.PetsAllowed != e.PetsAllowed:
			continue
		default:
			*result = append(*result, e)
		}
	}
	return result
}

type FilterEstateController struct{}

func (f FilterEstateController) Dispatch(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		f.Get(w, r)
	case "POST":
		f.Post(w, r)
	default:
		raiseMethodNotAllowed(w, r)
	}
}

func (f FilterEstateController) Post(w http.ResponseWriter, r *http.Request) {

	var filter FilterForm

	if err := filter.addBody(r.Body); err != nil {
		log.Println(r.URL, r.Method, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	estates := []models.Estate{
		{
			Id:            1,
			Name:          "Big house",
			AddressNumber: "222",
			AtFloor:       0,
			Floors:        2,
			Balconies:     4,
			Bedrooms:      7,
			Bathrooms:     5,
			Parking:       4,
			Description:   "A big house really nice",
			PetsAllowed:   1,
		},
		{
			Id:            2,
			Name:          "Small house",
			AddressNumber: "19",
			AtFloor:       0,
			Floors:        1,
			Balconies:     0,
			Bedrooms:      2,
			Bathrooms:     1,
			Parking:       0,
			Description:   "A small house",
			PetsAllowed:   1,
		},
		{
			Id:            3,
			Name:          "Apartment",
			AddressNumber: "202",
			AtFloor:       2,
			Floors:        1,
			Balconies:     1,
			Bedrooms:      2,
			Bathrooms:     2,
			Parking:       1,
			Description:   "A two bedroom apartment",
			PetsAllowed:   2,
		},
		{
			Id:            4,
			Name:          "Apartment B",
			AddressNumber: "604",
			AtFloor:       6,
			Floors:        1,
			Balconies:     1,
			Bedrooms:      3,
			Bathrooms:     2,
			Parking:       1,
			Description:   "A three bedroom apartment",
			PetsAllowed:   2,
		},
		{
			Id:            5,
			Name:          "Apartment C",
			AddressNumber: "403",
			AtFloor:       4,
			Floors:        1,
			Balconies:     1,
			Bedrooms:      2,
			Bathrooms:     2,
			Parking:       1,
			Description:   "Another two bedroom apartment",
			PetsAllowed:   1,
		},
		{
			Id:            6,
			Name:          "Apartment D",
			AddressNumber: "404",
			AtFloor:       4,
			Floors:        1,
			Balconies:     1,
			Bedrooms:      2,
			Bathrooms:     2,
			Parking:       1,
			Description:   "Another two bedroom apartment",
			PetsAllowed:   1,
		},
		{
			Id:            7,
			Name:          "House",
			AddressNumber: "333",
			AtFloor:       0,
			Floors:        1,
			Balconies:     0,
			Bedrooms:      2,
			Bathrooms:     2,
			Parking:       1,
			Description:   "Another house",
			PetsAllowed:   1,
		},
	}

	tmpl := template.Must(template.ParseFiles("views/estates-table.html"))
	tmpl.Execute(w, filter.filter(&estates))
}

func (f FilterEstateController) Get(w http.ResponseWriter, r *http.Request) {
	raiseMethodNotAllowed(w, r)
}
