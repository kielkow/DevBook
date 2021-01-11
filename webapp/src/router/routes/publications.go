package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesPublications = []Route{
	{
		URI:                   "/publications",
		Method:                http.MethodPost,
		Function:              controllers.CreatePublication,
		RequireAuthentication: true,
	},
}
