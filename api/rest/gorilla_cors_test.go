package rest

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
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

	// go http.ListenAndServe("127.0.0.1:8687", handler.Router)

	checkOptionRouteCors(t, handler.Router, realmResource, AnyOrigin, "*", "OPTIONS,GET" )
	checkOptionRouteCors(t, handler.Router, userResourceRoot, AnyOrigin, "*", "OPTIONS,GET,POST" )

        userById := "/api/user/123/"
	checkOptionRouteCors(t, handler.Router, userById, AnyOrigin, "*", "OPTIONS,GET,PUT,DELETE" )
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

	checkOptionRouteCors(t, handler.Router, "service1" + objectResource, internalSubNet, "*", "OPTIONS,GET,POST" )
	//checkOptionRouteCors(t, handler.Router, "service2" + classRootResource, internalSubNet, "*", "OPTIONS,GET,POST" )

	//classById := "/api/class/356/"
	//checkOptionRouteCors(t, handler.Router, "service2" + classById, internalSubNet, "*", "OPTIONS,DELETE" )
}

func checkOptionRouteCors(t *testing.T, router *mux.Router, requestPath string, allowedOrigin string, allowedHeader string, allowedMethods string) {
	request := http.Request{URL: &url.URL{Scheme: "http", Host: "127.0.0.1:8687", Path: requestPath},
		                Method: "OPTIONS"}
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, &request)
	assert.Equal(t, allowedOrigin, writer.Header().Get(AccessControlAllowOriginHeader))
	assert.Equal(t, allowedHeader, writer.Header().Get(AccessControlAllowHeadersHeader))
	assert.Equal(t, allowedMethods, writer.Header().Get(AccessControlAllowMethodsHeader))
}
