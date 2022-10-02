package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Israel-Ferreira/todo-s3/pkg/data"
	"github.com/Israel-Ferreira/todo-s3/pkg/services"
	"github.com/go-chi/chi/v5"
)

type UserControllerImpl struct {
	service services.UserService
}

func (uc *UserControllerImpl) GetAll(rw http.ResponseWriter, r *http.Request) {
	users, err := uc.service.GetAll()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(rw).Encode(&users); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (uc *UserControllerImpl) GetById(rw http.ResponseWriter, r *http.Request) {
	reqId := chi.URLParam(r, "id")

	uintId, err := strconv.ParseUint(reqId, 10, 64)

	if err != nil {
		log.Printf("Erro ao fazer a requisição: %v \n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := uc.service.GetById(uintId)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(rw).Encode(&user); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (uc *UserControllerImpl) DeleteById(rw http.ResponseWriter, r *http.Request) {
	reqId := chi.URLParam(r, "id")

	uintId, err := strconv.ParseUint(reqId, 10, 64)

	if err != nil {
		log.Printf("Erro ao fazer a requisição: %v \n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = uc.service.DeleteById(uintId)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func (uc *UserControllerImpl) Update(rw http.ResponseWriter, r *http.Request) {
	reqId := chi.URLParam(r, "id")

	uintId, err := strconv.ParseUint(reqId, 10, 64)

	if err != nil {
		log.Printf("Erro ao fazer a requisição: %v \n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(uintId)
}

func (uc *UserControllerImpl) Create(rw http.ResponseWriter, r *http.Request) {
	var userData data.UserData

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		log.Printf("Erro ao fazer a requisição: %v \n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	insertedId, err := uc.service.Create(userData)

	if err != nil {
		log.Printf("Erro ao fazer a requisição: %v \n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	createUserResp := &data.NewUserData{InsertedId: insertedId}

	rw.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(rw).Encode(createUserResp); err != nil {
		log.Printf("Erro ao fazer a requisição: %v \n", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (uc *UserControllerImpl) UploadProfilePhoto(rw http.ResponseWriter, r *http.Request) {

}
