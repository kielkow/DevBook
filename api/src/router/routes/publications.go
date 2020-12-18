package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPublications = []Route{
	{
		URI:                   "/publications",
		Method:                http.MethodPost,
		Function:              controllers.CreatePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications",
		Method:                http.MethodGet,
		Function:              controllers.SearchPublications,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodGet,
		Function:              controllers.SearchPublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "users/{userId}/publications",
		Method:                http.MethodGet,
		Function:              controllers.SearchUserPublications,
		RequireAuthentication: true,
	},
}
