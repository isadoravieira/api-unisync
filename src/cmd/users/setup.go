package usersroutes

import (
	"github.com/gorilla/mux"
)

// Setup places all routes inside the router
func Setup(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
