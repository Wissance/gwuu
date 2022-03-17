package rest

import (
	"github.com/gorilla/mux"
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
	userResourceRoot := "/api/user/"
	handler.HandleFunc(userResourceRoot, func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")
	handler.HandleFunc(userResourceRoot, func(writer http.ResponseWriter, request *http.Request) {

	}, "POST")
	userResourceById := "/api/user/{id:[0-9]+}/"
	handler.HandleFunc(userResourceById, func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")
	handler.HandleFunc(userResourceById, func(writer http.ResponseWriter, request *http.Request) {

	}, "PUT")
	handler.HandleFunc(userResourceById, func(writer http.ResponseWriter, request *http.Request) {

	}, "DELETE")

	realmOptionRouteName := stringFormatter.Format("{0}_{1}", realmResource, optionsRouteSuffix)
	route := handler.Router.Get(realmOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, realmResource, "*", "*", "OPTIONS,GET" )

	userRootOptionRouteName := stringFormatter.Format("{0}_{1}", userResourceRoot, optionsRouteSuffix)
	route = handler.Router.Get(userRootOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, userResourceRoot, "*", "*", "OPTIONS,GET,POST" )

	userByIdOptionRouteName := stringFormatter.Format("{0}_{1}", userResourceById, optionsRouteSuffix)
	route = handler.Router.Get(userByIdOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, userResourceById, "*", "*", "OPTIONS,GET,PUT,DELETE" )
}

func checkOptionRouteCors(t *testing.T, route *mux.Route, requestPath string, allowedOrigin string, allowedHeader string, allowedMethods string) {
	m, _ := route.Methods().GetMethods()
	assert.Equal(t, 1, len(m))
	assert.Equal(t, "OPTIONS", m[0])
	request := http.Request{URL: &url.URL{Scheme: "http", Host: "127.0.0.1", Path: requestPath},
		Method: "OPTIONS"}
	writer := httptest.NewRecorder()
	route.GetHandler().ServeHTTP(writer, &request)
	assert.Equal(t, allowedOrigin, writer.Header().Get(AccessControlAllowOriginHeader))
	assert.Equal(t, allowedHeader, writer.Header().Get(AccessControlAllowHeadersHeader))
	assert.Equal(t, allowedMethods, writer.Header().Get(AccessControlAllowMethodsHeader))
}
