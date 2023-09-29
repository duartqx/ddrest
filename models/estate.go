package models

type Estate struct {
	Id              int    `json:"id"`
	Name            string `json:"estate_name"`
	EstateType_Id   int    `json:"estate_type_id"`
	EstateStatus_Id int    `json:"estate_status_id"`
	Road_Id         int    `json:"estate_road_id"`
	AddressNumber   string `json:"estate_address_number"`
	AtFloor         int    `json:"estate_at_floor"`
	Floors          int    `json:"estate_floors"`
	Balconies       int    `json:"estate_balconies"`
	Bedrooms        int    `json:"estate_bedrooms"`
	Bathrooms       int    `json:"estate_bathrooms"`
	Parking         int    `json:"estate_parking"`
	PetsAllowed     bool   `json:"estate_pets_allowed"`
	Description     string `json:"estate_description"`
}
