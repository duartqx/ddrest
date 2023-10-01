package filters

import (
	"ddrest/models"
	"encoding/json"
	"io"
)

type FilterRealEstateForm struct {
	Name        string `json:"estate_name"`
	Floors      int    `json:"estate_floors"`
	Balconies   int    `json:"estate_balconies"`
	Bedrooms    int    `json:"estate_bedrooms"`
	Bathrooms   int    `json:"estate_bathrooms"`
	Parking     int    `json:"estate_parking"`
	PetsAllowed int    `json:"estate_pets_allowed"`
}

func (f *FilterRealEstateForm) AddBody(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(&f)
}

func (f FilterRealEstateForm) Filter() (results *[]models.RealEstate) {
	q := &models.Query{Values: []any{}}

	if f.Name != "" {
		q.Str += " Name LIKE ?"
		q.Values = append(q.Values, "%"+f.Name+"%")
	}

	if f.Floors > 0 {
		q.Add(" Floors >= ?", f.Floors)
	}
	if f.Balconies > 0 {
		q.Add(" Balconies >= ?", f.Balconies)
	}
	if f.Bedrooms > 0 {
		q.Add(" Bedrooms >= ?", f.Bedrooms)
	}
	if f.Parking > 0 {
		q.Add(" Parking >= ?", f.Parking)
	}
	if f.Bathrooms > 0 {
		q.Add(" Bathrooms >= ?", f.Bathrooms)
	}
	if f.PetsAllowed != 2 {
		q.Add(" PetsAllowed = ?", f.PetsAllowed)
	}

	if q.Str != "" {
		results = models.RealEstate{}.Filter(q)
	} else {
		results = models.RealEstate{}.All()
	}

	return results
}
