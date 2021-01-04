package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routeHomePage = Route{
	URI:                   "/home",
	Method:                http.MethodGet,
	Function:              controllers.RenderHomePage,
	RequireAuthentication: true,
}
