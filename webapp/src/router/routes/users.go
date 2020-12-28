package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var routesUser = []Route{
	{
		URI:                   "/create-user",
		Method:                http.MethodGet,
		Function:              controllers.RenderSignupScreen,
		RequireAuthentication: false,
	},
}
