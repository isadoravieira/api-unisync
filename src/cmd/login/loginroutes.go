package login

import (
	"net/http"

	"github.com/isadoravieira/api-unisync/src/cmd/login/handler"
	"github.com/isadoravieira/api-unisync/src/internal/domain/entity"
)

var LoginRoute = entity.Route{
	URI: "/login",
	Method: http.MethodPost,
	Function: handler.Login,
	AuthorizationRequired: false,
}