package filters

import (
	"ddrest/models"
	"encoding/json"
	"io"
	"strings"
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

func (f FilterRealEstateForm) Filter(estates *[]models.RealEstate) (results *[]models.RealEstate) {
	results = &[]models.RealEstate{}

	for _, e := range *estates {
		switch {
		case f.Name != "" && !strings.Contains(e.Name, f.Name):
			continue
		case f.Floors > e.Floors:
			continue
		case f.Balconies > e.Balconies:
			continue
		case f.Bedrooms > e.Bedrooms:
			continue
		case f.Parking > e.Parking:
			continue
		case f.Bathrooms > e.Bathrooms:
			continue
		case f.PetsAllowed != 0 && f.PetsAllowed != e.PetsAllowed:
			continue
		default:
			*results = append(*results, e)
		}
	}
	return results
}
