package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates HTTPs routes with its handlers
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	// Expvar exposes variables via HTTP in JSON format.
	router.Handle("/debug/vars", http.DefaultServeMux)

	return router
}
