package main

import "net/http"

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
		Index,
	},
	Route{
		"CommandList",
		"GET",
		"/commands",
		CommandList,
	},
	Route{
		"NetworkList",
		"GET",
		"/networks",
		NetworkList,
	},
	Route{
		"NetworkCreate",
		"POST",
		"/networks",
		NetworkCreate,
	},
	Route{
		"NetworkShow",
		"GET",
		"/networks/{name}",
		NetworkShow,
	},
}
