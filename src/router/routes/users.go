package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.SearchUsers,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.SearchUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/unfollow",
		Method:      http.MethodPost,
		Function:    controllers.UnfollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/followers",
		Method:      http.MethodGet,
		Function:    controllers.FindFollowers,
		RequireAuth: true,
	},
}
