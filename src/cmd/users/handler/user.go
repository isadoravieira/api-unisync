package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/isadoravieira/api-unisync/src/config"
	"github.com/isadoravieira/api-unisync/src/internal/domain/entity"
	"github.com/isadoravieira/api-unisync/src/internal/infraestructure/repository"
	"github.com/isadoravieira/api-unisync/src/pkg/responses"
)

func StoreUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.DomainError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user entity.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.DomainError(w, http.StatusBadRequest, err)
		return
	}

	db, err := config.Connect()
	if err != nil {
		responses.DomainError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	newUser, err := userRepo.Create(user)
	if err != nil {
		responses.DomainError(w, http.StatusInternalServerError, err)
		return
	}

	responses.DomainJSON(w, http.StatusCreated, newUser)
}

func IndexUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List user"))
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Query user by ID"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
