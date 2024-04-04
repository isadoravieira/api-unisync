package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/isadoravieira/api-unisync/src/config"
	"github.com/isadoravieira/api-unisync/src/internal/domain/entity"
	"github.com/isadoravieira/api-unisync/src/internal/infraestructure/repository"
)

func StoreUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user entity.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := config.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userRepo.Create(user)
	
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
