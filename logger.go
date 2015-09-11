package main

import (
	"log"
	"net/http"
	"time"

	"github.com/juandiegoh/jaydiapps/expti"
)

// Logger decorator of http.Handler
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		expti.AddCount(r.URL.RequestURI())

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
