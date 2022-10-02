package controllers

import "net/http"

type UserController interface {
	CrudController
	UploadProfilePhoto(rw http.ResponseWriter, r *http.Request)
}
