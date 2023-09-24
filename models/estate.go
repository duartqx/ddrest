package models

import (
	"encoding/json"
	"strconv"
)

type t interface{}

type Estate struct {
	Id            int    `json:"id"`
	Name          string `json:"estate_name"`
	RoadId        int    `json:"estate_road_id"`
	AddressNumber string `json:"estate_address_number"`
	EstateTypeId  int    `json:"estate_type_id"`
	AtFloor       int    `json:"estate_at_floor"`
	Floors        int    `json:"estate_floors"`
	Balconies     int    `json:"estate_balconies"`
	Bedrooms      int    `json:"estate_bedrooms"`
	Bathrooms     int    `json:"estate_bathrooms"`
	Parking       int    `json:"estate_parking"`
	PetsAllowed   bool   `json:"estate_pets_allowed"`
	Description   string `json:"estate_description"`
	StatusId      int    `json:"estate_status_id"`
}

func (e *Estate) Set(field string, value interface{}) {
	switch field {
	case "Id":
		e.SetId(value.(int))
	case "Name":
		e.SetName(value.(string))
	case "RoadId":
		e.SetRoadId(value.(int))
	case "AddressNumber":
		e.SetAddressNumber(value.(string))
	case "EstateTypeId":
		e.SetEstateTypeId(value.(int))
	case "AtFloor":
		e.SetAtFloor(value.(int))
	case "Floors":
		e.SetFloors(value.(int))
	case "Balconies":
		e.SetBalconies(value.(int))
	case "Bedrooms":
		e.SetBedrooms(value.(int))
	case "Bathrooms":
		e.SetBathrooms(value.(int))
	case "Parking":
		e.SetParking(value.(int))
	case "PetsAllowed":
		e.SetPetsAllowed(value.(bool))
	case "Description":
		e.SetDescription(value.(string))
	case "StatusId":
		e.SetStatusId(value.(int))
	}
}

func (e *Estate) SetId(id int) *Estate {
	e.Id = id
	return e
}

func (e *Estate) SetName(name string) *Estate {
	e.Name = name
	return e
}

func (e *Estate) SetRoadId(id int) *Estate {
	e.RoadId = id
	return e
}

func (e *Estate) SetAddressNumber(number string) *Estate {
	e.AddressNumber = number
	return e
}

func (e *Estate) SetEstateTypeId(id int) *Estate {
	e.EstateTypeId = id
	return e
}

func (e *Estate) SetAtFloor(floor int) *Estate {
	e.AtFloor = floor
	return e
}

func (e *Estate) SetFloors(floors int) *Estate {
	e.Floors = floors
	return e
}

func (e *Estate) SetBalconies(balconies int) *Estate {
	e.Balconies = balconies
	return e
}

func (e *Estate) SetBedrooms(bedrooms int) *Estate {
	e.Bedrooms = bedrooms
	return e
}

func (e *Estate) SetBathrooms(bathrooms int) *Estate {
	e.Bathrooms = bathrooms
	return e
}

func (e *Estate) SetParking(parkings int) *Estate {
	e.Parking = parkings
	return e
}

func (e *Estate) SetPetsAllowed(allowed bool) *Estate {
	e.PetsAllowed = allowed
	return e
}

func (e *Estate) SetDescription(description string) *Estate {
	e.Description = description
	return e
}

func (e *Estate) SetStatusId(id int) *Estate {
	e.StatusId = id
	return e
}

func (e *Estate) Get(field string, value t) error {
	return nil
}

func (e *Estate) UnmarshalJSON(data []byte) error {
	type Alias Estate

	var jsonData map[string]*json.RawMessage

	if err := json.Unmarshal(data, &jsonData); err != nil {
		return err
	}

	stringAttributes := []map[string]string{
		{"attr": "Name", "jsonKey": "estate_name"},
		{"attr": "AddressNumber", "jsonKey": "estate_address_number"},
		{"attr": "Description", "jsonKey": "estate_description"},
	}

	for _, field := range stringAttributes {
		if value, ok := jsonData[field["jsonKey"]]; ok {
			var unmsh string
			if err := json.Unmarshal(*value, &unmsh); err != nil {
				return err
			}
			e.Set(field["attr"], unmsh)
		}
	}

	intAttributes := []map[string]string{
		{"attr": "RoadId", "jsonKey": "estate_road_id"},
		{"attr": "EstateTypeId", "jsonKey": "estate_type_id"},
		{"attr": "AtFloor", "jsonKey": "estate_at_floor"},
		{"attr": "Balconies", "jsonKey": "estate_balconies"},
		{"attr": "Bedrooms", "jsonKey": "estate_bedrooms"},
		{"attr": "Bathrooms", "jsonKey": "estate_bathrooms"},
		{"attr": "Parking", "jsonKey": "estate_parking"},
		{"attr": "StatusId", "jsonKey": "estate_status_id"},
	}

	for _, field := range intAttributes {
		if value, ok := jsonData[field["jsonKey"]]; ok {
			var unmsh string
			if err := json.Unmarshal(*value, &unmsh); err != nil {
				return err
			}
			intValue, err := strconv.Atoi(unmsh)
			if err != nil {
				return err
			}
			e.Set(field["attr"], intValue)
		}
	}

	boolAttributes := []map[string]string{
		{"attr": "PetsAllowed", "jsonKey": "estate_pets_allowed"},
	}

	for _, field := range boolAttributes {
		if value, ok := jsonData[field["jsonKey"]]; ok {
			var unmsh string
			if err := json.Unmarshal(*value, &unmsh); err != nil {
				return err
			}
			e.Set(field["attr"], bool(unmsh == "1"))
		}
	}

	return nil
}
