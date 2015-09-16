package main

import (
	"net/http"

	"github.com/juandiegoh/jaydiapps/expti"
)

// Reporter go_expvar reporter wrapper for HTTPHandler
func Reporter(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expti.AddCount("routes_" + name + "_" + r.URL.RequestURI())
		inner.ServeHTTP(w, r)
	})
}
