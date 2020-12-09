package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.SearchUsers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		Function:              controllers.SearchUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
}
