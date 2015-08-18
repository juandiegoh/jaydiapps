package main

import "net/http"

// Route specifies an HTTP Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array of Route
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"LoveIndex",
		"GET",
		"/loves",
		LoveIndex,
	},
	Route{
		"LoveShow",
		"GET",
		"/loves/{id}",
		LoveShow,
	},
}
