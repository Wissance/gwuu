package rest

import (
	"github.com/stretchr/testify/assert"
	"github.com/wissance/stringFormatter"
	"net/http"
	"testing"
)

func TestHandleFuncWithCorsWithAnyOrigin(t *testing.T) {
	// Assign routes for resources with GET only and FULL CRUD
	handler := NewWebApiHandler(true, AnyOrigin)
	// Get only method
	realmResource := "/api/realm/"
	handler.HandleFunc(realmResource, func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")
	// full crud
	handler.HandleFunc("/api/user/", func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")
	handler.HandleFunc("/api/user/", func(writer http.ResponseWriter, request *http.Request) {

	}, "POST")
	handler.HandleFunc("/api/user/{id:[0-9]+}/", func(writer http.ResponseWriter, request *http.Request) {

	}, "PUT")
	handler.HandleFunc("/api/user/{id:[0-9]+}/", func(writer http.ResponseWriter, request *http.Request) {

	}, "DELETE")

	realmOptionRoute := stringFormatter.Format("{0}_{1}", realmResource, optionsRouteSuffix)
	route := handler.Router.Get(realmOptionRoute)
	assert.NotNil(t, route)
	assert.Equal(t, "OPTIONS", route.Methods())
}
