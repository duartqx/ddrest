package models

type RealEstate struct {
	Id            int    `json:"id"`
	Name          string `json:"estate_name"`
	EstateType    int    `json:"estate_type_id"`
	EstateStatus  int    `json:"estate_status_id"`
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

func (re RealEstate) EstateTypeChoices() *map[int]string {
	return &map[int]string{
		0: "",
		1: "House",
		2: "Apartment",
		3: "Apartment Building",
		4: "Room",
		5: "Studio",
		6: "Store",
		7: "Area",
		8: "Lot",
	}
}

func (re RealEstate) EstateStatusChoices() *map[int]string {
	return &map[int]string{
		0: "",
		1: "Empty",
		2: "Rented",
		3: "New",
		4: "Needs repairs",
		5: "Irregular",
	}
}

func (re RealEstate) PetsAllowedChoices() *map[int]string {
	return &map[int]string{
		0: "No",
		1: "Yes",
		2: "",
	}
}
