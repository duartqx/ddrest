package controllers

import "net/http"

type IController interface {
	Dispatch(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	Endpoint string
}
