package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesLogout = Route{
	URI:                   "/logout",
	Method:                http.MethodGet,
	Function:              controllers.Signout,
	RequireAuthentication: true,
}
