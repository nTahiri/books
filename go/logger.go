/*
 * Books Store
 *
 * A sample API that uses a Books store as an example to demonstrate features in the OpenAPI 3.0 specification
 *
 * API version: 0.0.1
 * Contact: nabil.tahri@polymtl.ca
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package book

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
