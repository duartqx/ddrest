package models

import (
	"database/sql"
	"fmt"
	"log"
)

type RealEstate struct {
	Id           int    `json:"id"`
	Name         string `json:"estate_name"`
	EstateType   int    `json:"estate_type_id"`
	EstateStatus int    `json:"estate_status_id"`
	// Road_Id       int    `json:"estate_road_id"`
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

func (re RealEstate) EstateStatusChoices() *map[int]string {
	return &map[int]string{
		0: "Empty",
		1: "Rented",
		2: "New",
		3: "Needs repairs",
		4: "Irregular",
	}
}

func (re RealEstate) PetsAllowedChoices() *map[int]string {
	return &map[int]string{
		0: "No",
		1: "Yes",
		2: "Null",
	}
}

func (re RealEstate) buildResults(rows *sql.Rows) *[]RealEstate {
	var (
		results []RealEstate
		tmp     RealEstate
	)

	for rows.Next() {
		if err := rows.Scan(
			&tmp.Id,
			&tmp.Name,
			&tmp.EstateType,
			&tmp.EstateStatus,
			&tmp.AddressNumber,
			&tmp.AtFloor,
			&tmp.Floors,
			&tmp.Balconies,
			&tmp.Bedrooms,
			&tmp.Bathrooms,
			&tmp.Parking,
			&tmp.PetsAllowed,
			&tmp.Description,
		); err != nil {
			log.Fatal(err)
		}
		results = append(results, tmp)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &results
}

func (re RealEstate) All() *[]RealEstate {

	rows, err := repo.db.Query(fmt.Sprintf("%s;", reMeta.baseSelect()))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return re.buildResults(rows)
}

func (re RealEstate) Filter(query *Query) *[]RealEstate {

	select_str := fmt.Sprintf("%s WHERE %s;", reMeta.baseSelect(), query.Str)

	rows, err := repo.db.Query(select_str, query.Values...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return re.buildResults(rows)
}

type realEstateMeta struct {
	tableName string
}

func (m realEstateMeta) baseSelect() string {
	return fmt.Sprintf(`
		SELECT 
			Id,
			Name,
			EstateType,
			EstateStatus,
			AddressNumber,
			AtFloor,
			Floors,
			Balconies,
			Bedrooms,
			Bathrooms,
			Parking,
			PetsAllowed,
			Description
		FROM %s
	`, m.tableName)
}

var reMeta realEstateMeta = realEstateMeta{
	tableName: "RealEstate",
}
