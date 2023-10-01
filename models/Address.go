package models

type Country struct {
	Id   int
	Name string
}

type State struct {
	Id         int
	Name       string
	Country_id int
}

type City struct {
	Id       int
	Name     string
	State_id int
}

type Road struct {
	Id      int
	Name    string
	Zip     string
	City_id int
}
