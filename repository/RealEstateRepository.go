package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"ddrest/models"
)

const realEstateTable string = "RealEstate"

type RealEstateRepository struct {
	tmp models.RealEstate
}

func (re RealEstateRepository) buildResults(rows *sql.Rows) *[]models.RealEstate {
	var (
		results []models.RealEstate
		tmp     models.RealEstate
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

func (re RealEstateRepository) Filter() *[]models.RealEstate {
	q := &Query{Values: []any{}}

	if re.tmp.Name != "" {
		q.Add(" Name LIKE ?", "%"+re.tmp.Name+"%")
	}
	if re.tmp.EstateType > 0 {
		q.Add(" EstateType = ?", re.tmp.EstateType)
	}
	if re.tmp.EstateStatus > 0 {
		q.Add(" EstateStatus = ?", re.tmp.EstateStatus)
	}
	if re.tmp.AddressNumber != "" {
		q.Add(" AddressNumber LIKE ?", "%"+re.tmp.AddressNumber+"%")
	}
	if re.tmp.AtFloor > 0 {
		q.Add(" AtFloor = ?", re.tmp.AtFloor)
	}
	if re.tmp.Floors > 0 {
		q.Add(" Floors >= ?", re.tmp.Floors)
	}
	if re.tmp.Balconies > 0 {
		q.Add(" Balconies >= ?", re.tmp.Balconies)
	}
	if re.tmp.Bedrooms > 0 {
		q.Add(" Bedrooms >= ?", re.tmp.Bedrooms)
	}
	if re.tmp.Parking > 0 {
		q.Add(" Parking >= ?", re.tmp.Parking)
	}
	if re.tmp.Bathrooms > 0 {
		q.Add(" Bathrooms >= ?", re.tmp.Bathrooms)
	}
	if re.tmp.Parking > 0 {
		q.Add(" Parking >= ?", re.tmp.Parking)
	}
	if re.tmp.PetsAllowed != 2 {
		q.Add(" PetsAllowed = ?", re.tmp.PetsAllowed)
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

func (re *RealEstateRepository) All() *[]models.RealEstate {

	rows, err := repo.db.Query(fmt.Sprintf("SELECT * FROM %s;", realEstateTable))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return re.buildResults(rows)
}

func (m *RealEstateRepository) Populate(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(&m.tmp)
}
