package entity

import "net/http"

// Route represents all api routes
type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	AuthorizationRequired bool
}
