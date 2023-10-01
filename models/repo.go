package models

import "database/sql"

type Repo struct {
	db *sql.DB
}

var repo Repo

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
