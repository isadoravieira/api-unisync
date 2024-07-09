package router

import (
	"github.com/gorilla/mux"
	"github.com/isadoravieira/api-unisync/src/cmd/login"
	usersroutes "github.com/isadoravieira/api-unisync/src/cmd/users"
)

// Setup places all routes inside the router
func Setup(r *mux.Router) *mux.Router {
	routes := usersroutes.UsersRoutes
	routes = append(routes, login.LoginRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
