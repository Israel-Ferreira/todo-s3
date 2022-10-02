package controllers

import "net/http"

type CrudController interface {
	GetAll(rw http.ResponseWriter, r *http.Request)
	GetById(rw http.ResponseWriter, r *http.Request)
	DeleteById(rw http.ResponseWriter, r *http.Request)
	Update(rw http.ResponseWriter, r *http.Request)
	Create(rw http.ResponseWriter, r *http.Request)
}
