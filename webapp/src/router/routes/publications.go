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
	{
		URI:                   "/publications/{publicationId}/like",
		Method:                http.MethodPost,
		Function:              controllers.LikePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}/dislike",
		Method:                http.MethodPost,
		Function:              controllers.DislikePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}/update",
		Method:                http.MethodGet,
		Function:              controllers.RenderUpdatePublicationPage,
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
}
