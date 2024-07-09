package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/isadoravieira/api-unisync/src/config"
	"github.com/isadoravieira/api-unisync/src/internal/domain/entity"
	"github.com/isadoravieira/api-unisync/src/internal/infraestructure/repository"
	"github.com/isadoravieira/api-unisync/src/pkg/responses"
	"github.com/isadoravieira/api-unisync/src/pkg/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NewUserRepository(db)
	userInDatabase, err := repo.QueryByEmail(user.Email)
	if err != nil {
		responses.DomainError(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(userInDatabase.Password, user.Password); err !=  nil {
		responses.DomainError(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte("You are in!!"))
}