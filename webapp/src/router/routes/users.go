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
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
}
