package filters

import (
	"ddrest/models"
	"encoding/json"
	"fmt"
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
	query := &models.Query{Values: []any{}}

	if f.Name != "" {
		query.Str += " Name LIKE ?"
		query.Values = append(query.Values, fmt.Sprintf("%%%s%%", f.Name))
	}
	if f.Floors > 0 {
		if query.Str != "" {
			query.Str += " AND Floors >= ?"
		} else {
			query.Str += " Floors >= ?"
		}
		query.Values = append(query.Values, f.Floors)
	}
	if f.Balconies > 0 {
		if query.Str != "" {
			query.Str += " AND Balconies >= ?"
		} else {
			query.Str += " Balconies >= ?"
		}
		query.Values = append(query.Values, f.Balconies)
	}
	if f.Bedrooms > 0 {
		if query.Str != "" {
			query.Str += " AND Bedrooms >= ?"
		} else {
			query.Str += " Bedrooms >= ?"
		}
		query.Values = append(query.Values, f.Bedrooms)
	}
	if f.Parking > 0 {
		if query.Str != "" {
			query.Str += " AND Parking >= ?"
		} else {
			query.Str += " Parking >= ?"
		}
		query.Values = append(query.Values, f.Parking)
	}
	if f.Bathrooms > 0 {
		if query.Str != "" {
			query.Str += " AND Bathrooms >= ?"
		} else {
			query.Str += " Bathrooms >= ?"
		}
		query.Values = append(query.Values, f.Bathrooms)
	}
	if f.PetsAllowed != 2 {
		if query.Str != "" {
			query.Str += " AND PetsAllowed = ?"
		} else {
			query.Str += " PetsAllowed = ?"
		}
		query.Values = append(query.Values, f.PetsAllowed)
	}

	if query.Str != "" {
		results = models.RealEstate{}.Filter(query)
	} else {
		results = models.RealEstate{}.All()
	}

	return results
}
