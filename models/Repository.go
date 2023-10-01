package models

import "database/sql"

type Repository struct {
	db *sql.DB
}

var repo Repository

func InitRepository(db *sql.DB) {
	repo.db = db
}

type Model interface {
	Filter(params map[string]any) *sql.Rows
	GetById(id int) *sql.Rows
}

type Query struct {
	Str    string
	Values []any
}

func (q Query) and() string {
	switch {
	case q.Str != "":
		return " AND"
	default:
		return ""
	}
}

func (q *Query) Add(query string, value any) {
	q.Str += q.and() + query
	q.Values = append(q.Values, value)
}
