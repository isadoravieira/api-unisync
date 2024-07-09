package usersroutes

import (
	"net/http"

	"github.com/isadoravieira/api-unisync/src/cmd/users/handler"
	"github.com/isadoravieira/api-unisync/src/internal/domain/entity"
)

var UsersRoutes = []entity.Route {
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: handler.StoreUser,
		AuthorizationRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: handler.IndexUser,
		AuthorizationRequired: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodGet,
		Function: handler.ShowUser,
		AuthorizationRequired: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodPut,
		Function: handler.UpdateUser,
		AuthorizationRequired: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete,
		Function: handler.DeleteUser,
		AuthorizationRequired: false,
	},
}