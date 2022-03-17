package rest

import (
	"github.com/stretchr/testify/assert"
	"github.com/wissance/stringFormatter"
	"net/http"
	"net/http/httptest"
	"net/url"
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
	m, _ := route.Methods().GetMethods()
	assert.Equal(t, 1, len(m))
	assert.Equal(t, "OPTIONS", m[0])
	request := http.Request{URL: &url.URL{Scheme: "http", Host: "127.0.0.1", Path: realmResource},
		                Method: "OPTIONS"}
	writer := httptest.NewRecorder()
	route.GetHandler().ServeHTTP(writer, &request)
	assert.Equal(t, "*", writer.Header().Get(AccessControlAllowOriginHeader))
	assert.Equal(t, "*", writer.Header().Get(AccessControlAllowHeadersHeader))
	assert.Equal(t, "OPTIONS,GET", writer.Header().Get(AccessControlAllowMethodsHeader))

}
