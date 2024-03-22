package router

import (
	"github.com/gorilla/mux"
	usersroutes "github.com/isadoravieira/api-unisync/src/cmd/users"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return usersroutes.Setup(r)
}