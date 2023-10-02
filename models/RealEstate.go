package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

const realEstateTable string = "RealEstate"

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

func (re RealEstate) buildResults(rows *sql.Rows) *[]RealEstate {
	var results []RealEstate

	for rows.Next() {
		if err := rows.Scan(
			&re.Id,
			&re.Name,
			&re.EstateType,
			&re.EstateStatus,
			&re.AddressNumber,
			&re.AtFloor,
			&re.Floors,
			&re.Balconies,
			&re.Bedrooms,
			&re.Bathrooms,
			&re.Parking,
			&re.PetsAllowed,
			&re.Description,
		); err != nil {
			log.Fatal(err)
		}
		results = append(results, re)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &results
}

func (re RealEstate) Filter() *[]RealEstate {
	q := &Query{Values: []any{}}

	if re.Name != "" {
		q.Add(" Name LIKE ?", "%"+re.Name+"%")
	}
	if re.EstateType > 0 {
		q.Add(" EstateType = ?", re.EstateType)
	}
	if re.EstateStatus > 0 {
		q.Add(" EstateStatus = ?", re.EstateStatus)
	}
	if re.AddressNumber != "" {
		q.Add(" AddressNumber LIKE ?", "%"+re.AddressNumber+"%")
	}
	if re.AtFloor > 0 {
		q.Add(" AtFloor = ?", re.AtFloor)
	}
	if re.Floors > 0 {
		q.Add(" Floors >= ?", re.Floors)
	}
	if re.Balconies > 0 {
		q.Add(" Balconies >= ?", re.Balconies)
	}
	if re.Bedrooms > 0 {
		q.Add(" Bedrooms >= ?", re.Bedrooms)
	}
	if re.Parking > 0 {
		q.Add(" Parking >= ?", re.Parking)
	}
	if re.Bathrooms > 0 {
		q.Add(" Bathrooms >= ?", re.Bathrooms)
	}
	if re.Parking > 0 {
		q.Add(" Parking >= ?", re.Parking)
	}
	if re.PetsAllowed != 2 {
		q.Add(" PetsAllowed = ?", re.PetsAllowed)
	}

	if q.Str == "" {
		return re.All()
	}

	select_str := fmt.Sprintf("SELECT * FROM %s WHERE %s;", realEstateTable, q.Str)
	rows, err := repo.db.Query(select_str, q.Values...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return re.buildResults(rows)
}

func (re RealEstate) All() *[]RealEstate {

	rows, err := repo.db.Query(fmt.Sprintf("SELECT * FROM %s;", realEstateTable))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return re.buildResults(rows)
}

func (re RealEstate) Create(realEstates *[]RealEstate) int {
	query := &Query{
		Str: `
			INSERT INTO RealEstate (
				Name,
				AddressNumber,
				AtFloor,
				Floors,
				Balconies,
				Bedrooms,
				Bathrooms,
				Parking,
				PetsAllowed,
				EstateType,
				EstateStatus,
				Description
			)
		VALUES
		`,
		Values: []any{},
	}

	for _, e := range *realEstates {
		query.Str += " (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
		query.Values = append(
			query.Values,
			e.Name,
			e.AddressNumber,
			e.AtFloor,
			e.Floors,
			e.Balconies,
			e.Bedrooms,
			e.Bathrooms,
			e.Parking,
			e.PetsAllowed,
			e.EstateType,
			e.EstateStatus,
			e.Description,
		)
	}

	query.Str = fmt.Sprintf("%s;", query.Str[:len(query.Str)-1])

	res, err := repo.db.Exec(query.Str, query.Values...)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return int(rows)
}

func (re *RealEstate) Populate(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(&re)
}
