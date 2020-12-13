/*
 * Books Store
 *
 * A sample API that uses a petstore as an example to demonstrate features in the OpenAPI 3.0 specification
 *
 * API version: 0.0.1
 * Contact: nabil.tahri@polymtl.ca
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/books/go"
)

func main() {
	log.Printf("Server started")

	DefaultApiService := openapi.NewDefaultApiService()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
