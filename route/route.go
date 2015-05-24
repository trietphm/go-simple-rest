package route

import (
	"go-simple-rest/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"Posts",
		"GET",
		"/posts",
		handlers.List,
	},
	Route{
		"Query",
		"GET",
		"/query",
		handlers.Query,
	},
	Route{
		"Distinc",
		"GET",
		"/posts/distinc",
		handlers.Distinc,
	},
	Route{
		"CreatePost",
		"POST",
		"/posts",
		handlers.Create,
	},
	Route{
		"GetPost",
		"GET",
		"/post/{id}",
		handlers.GetPostById,
	},
	Route{
		"UpdatePost",
		"PUT",
		"/posts",
		handlers.UpdatePost,
	},
	Route{
		"DeletePost",
		"DELETE",
		"/post/{id}",
		handlers.DeletePost,
	},
}
