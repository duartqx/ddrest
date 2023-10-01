package models

type Estate struct {
	Id            int    `json:"id"`
	Name          string `json:"estate_name"`
	EstateType    int    `json:"estate_type_id"`
	EstateStatus  int    `json:"estate_status_id"`
	Road_Id       int    `json:"estate_road_id"`
	AddressNumber string `json:"estate_address_number"`
	AtFloor       int    `json:"estate_at_floor"`
	Floors        int    `json:"estate_floors"`
	Balconies     int    `json:"estate_balconies"`
	Bedrooms      int    `json:"estate_bedrooms"`
	Bathrooms     int    `json:"estate_bathrooms"`
	Parking       int    `json:"estate_parking"`
	PetsAllowed   int    `json:"estate_pets_allowed"`
	Description   string `json:"estate_description"`
}

func (e Estate) EstateTypeChoices() *map[int]string {
	return &map[int]string{
		0: "House",
		1: "Apartment",
		2: "Apartment Building",
		3: "Room",
		4: "Studio",
		5: "Store",
		6: "Area",
		7: "Lot",
	}
}

func (e Estate) EstateStatusChoices() *map[int]string {
	return &map[int]string{
		0: "Empty",
		1: "Rented",
		2: "New",
		3: "Needs repairs",
		4: "Irregular",
	}
}

func (e Estate) PetsAllowedChoices() *map[int]string {
	return &map[int]string{
		0: "Null",
		1: "Yes",
		2: "No",
	}
}
