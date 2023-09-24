package models

import "database/sql"

type Model interface {
	Filter(params map[string]any) *sql.Rows
	GetById(id int) *sql.Rows
}
