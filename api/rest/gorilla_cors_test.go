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
	handler.HandleFunc(handler.Router, realmResource, func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")
	// full crud
	userResourceRoot := "/api/user/"
	handler.HandleFunc(handler.Router, userResourceRoot, func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")
	handler.HandleFunc(handler.Router, userResourceRoot, func(writer http.ResponseWriter, request *http.Request) {

	}, "POST")
	userResourceById := "/api/user/{id:[0-9]+}/"
	handler.HandleFunc(handler.Router, userResourceById, func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")
	handler.HandleFunc(handler.Router, userResourceById, func(writer http.ResponseWriter, request *http.Request) {

	}, "PUT")
	handler.HandleFunc(handler.Router, userResourceById, func(writer http.ResponseWriter, request *http.Request) {

	}, "DELETE")

	realmOptionRouteName := stringFormatter.Format("{0}_{1}", realmResource, optionsRouteSuffix)
	route := handler.Router.Get(realmOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, realmResource, AnyOrigin, "*", "OPTIONS,GET" )

	userRootOptionRouteName := stringFormatter.Format("{0}_{1}", userResourceRoot, optionsRouteSuffix)
	route = handler.Router.Get(userRootOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, userResourceRoot, AnyOrigin, "*", "OPTIONS,GET,POST" )

	userByIdOptionRouteName := stringFormatter.Format("{0}_{1}", userResourceById, optionsRouteSuffix)
	route = handler.Router.Get(userByIdOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, userResourceById, AnyOrigin, "*", "OPTIONS,GET,PUT,DELETE" )
}

func TestHandleFuncForSubRouterAndSpecificOrigin(t *testing.T) {
	// there is no sub router access yet ...
	internalSubNet := "192.168.30.0"
	handler := NewWebApiHandler(true, internalSubNet)
	service1Router := handler.Router.PathPrefix("service1").Subrouter()

	objectResource := "/api/object/"
	handler.HandleFunc(service1Router, objectResource, func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")

	handler.HandleFunc(service1Router, objectResource, func(writer http.ResponseWriter, request *http.Request) {

	}, "POST")

	service2Router := handler.Router.PathPrefix("service2").Subrouter()
	classRootResource := "/api/class/"

	handler.HandleFunc(service2Router, classRootResource, func(writer http.ResponseWriter, request *http.Request) {

	}, "GET")

	handler.HandleFunc(service2Router, classRootResource, func(writer http.ResponseWriter, request *http.Request) {

	}, "POST")

	classByIdResource := "/api/class/{id}/"
	handler.HandleFunc(service2Router, classByIdResource, func(writer http.ResponseWriter, request *http.Request) {

	}, "DELETE")

	objectOptionRouteName := stringFormatter.Format("{0}_{1}", objectResource, optionsRouteSuffix)
	route := service1Router.Get(objectOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, objectResource, internalSubNet, "*", "OPTIONS,GET,POST" )

	classRootOptionRouteName := stringFormatter.Format("{0}_{1}", classRootResource, optionsRouteSuffix)
	route = service2Router.Get(classRootOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, classRootResource, internalSubNet, "*", "OPTIONS,GET,POST" )

	classByIdOptionRouteName := stringFormatter.Format("{0}_{1}", classByIdResource, optionsRouteSuffix)
	route = service2Router.Get(classByIdOptionRouteName)
	assert.NotNil(t, route)
	checkOptionRouteCors(t, route, classByIdResource, internalSubNet, "*", "OPTIONS,DELETE" )
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
