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
	{
		URI:                   "/search-users",
		Method:                http.MethodGet,
		Function:              controllers.RenderUsersPage,
		RequireAuthentication: true,
	},
	{
		URI:                   "/user/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.RenderUserProfile,
		RequireAuthentication: true,
	},
}
