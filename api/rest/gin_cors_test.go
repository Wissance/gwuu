package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wissance/gwuu/testingutils"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGinHandleFuncWithCorsWithAnyOrigin(t *testing.T) {
	handler := NewGinBasedWebApiHandler(true, AnyOrigin)
	handler.Router.RedirectTrailingSlash = true

	realmResource := "/api/realm"
	realmRoutes := handler.Router.Group(realmResource)
	handler.GET(realmRoutes, "/", func(ctx *gin.Context) {})

	userResource := "/api/user"
	userResourceByIdPath := "/:id/"
	userRoutes := handler.Router.Group(userResource)
	handler.GET(userRoutes, "/", func(ctx *gin.Context) {})
	handler.POST(userRoutes, "/", func(ctx *gin.Context) {})

	handler.GET(userRoutes, userResourceByIdPath, func(ctx *gin.Context) {})
	handler.PUT(userRoutes, userResourceByIdPath, func(ctx *gin.Context) {})
	handler.DELETE(userRoutes, userResourceByIdPath, func(ctx *gin.Context) {})

	// Requests with trailing slashes ...
	checkGinOptionRouteCors(t, handler.Router, realmResource+"/", AnyOrigin, "*", "OPTIONS,GET")
	checkGinOptionRouteCors(t, handler.Router, userResource+"/", AnyOrigin, "*", "OPTIONS,GET,POST")

	checkGinRouteCors(t, handler.Router, "GET", realmResource+"/", AnyOrigin)

	checkGinRouteCors(t, handler.Router, "GET", userResource+"/", AnyOrigin)
	checkGinRouteCors(t, handler.Router, "POST", userResource+"/", AnyOrigin)

	userById := "/api/user/123/"
	checkGinOptionRouteCors(t, handler.Router, userById, AnyOrigin, "*", "OPTIONS,GET,PUT,DELETE")

	checkGinRouteCors(t, handler.Router, "GET", userById, AnyOrigin)
	checkGinRouteCors(t, handler.Router, "PUT", userById, AnyOrigin)
	checkGinRouteCors(t, handler.Router, "DELETE", userById, AnyOrigin)
}

func TestGinHandleFuncForSubRouterAndSpecificOrigin(t *testing.T) {
	// there is no sub router access yet ...
	internalSubNet := "192.168.30.0"
	handler := NewGinBasedWebApiHandler(true, internalSubNet)
	service1Router := handler.Router.Group("/service1/")

	objectResource := "/api/object/"
	handler.GET(service1Router, objectResource, func(ctx *gin.Context) {})

	handler.POST(service1Router, objectResource, func(ctx *gin.Context) {})

	service2Router := handler.Router.Group("/service2/")
	classRootResource := "/api/class/"

	handler.GET(service2Router, classRootResource, func(ctx *gin.Context) {})

	handler.POST(service2Router, classRootResource, func(ctx *gin.Context) {})

	classByIdResource := "/api/class/:id/"
	handler.DELETE(service2Router, classByIdResource, func(ctx *gin.Context) {})

	checkGinOptionRouteCors(t, handler.Router, "/service1"+objectResource, internalSubNet, "*", "OPTIONS,GET,POST")
	checkGinOptionRouteCors(t, handler.Router, "/service2"+classRootResource, internalSubNet, "*", "OPTIONS,GET,POST")

	checkGinRouteCors(t, handler.Router, "GET", "/service1"+objectResource, internalSubNet)
	checkGinRouteCors(t, handler.Router, "POST", "/service1"+objectResource, internalSubNet)

	checkGinRouteCors(t, handler.Router, "GET", "/service2"+classRootResource, internalSubNet)
	checkGinRouteCors(t, handler.Router, "POST", "/service2"+classRootResource, internalSubNet)

	classById := "/api/class/356/"
	checkGinOptionRouteCors(t, handler.Router, "/service2"+classById, internalSubNet, "*", "OPTIONS,DELETE")

	checkGinRouteCors(t, handler.Router, "DELETE", "/service2"+classById, internalSubNet)
}

/*func TestGinHandleFuncForSubRouterSameName(t *testing.T) {
	internalSubNet := "192.168.30.0"
	handler := NewGinBasedWebApiHandler(true, internalSubNet)
	objectResource := "/api/object/"
	handler.HandleFunc(handler.Router, objectResource, func(writer http.ResponseWriter, request *http.Request) {
	}, "GET")
	service1Router := handler.Router.PathPrefix("/service1").Subrouter()
	handler.HandleFunc(service1Router, objectResource, func(writer http.ResponseWriter, request *http.Request) {
	}, "POST")
	checkMuxOptionRouteCors(t, handler.Router, "/api/object/", internalSubNet, "*", "OPTIONS,GET")
	checkMuxOptionRouteCors(t, handler.Router, "/service1/api/object/", internalSubNet, "*", "OPTIONS,POST")
}*/

func checkGinOptionRouteCors(t *testing.T, router *gin.Engine, requestPath string, allowedOrigin string,
	allowedHeader string, allowedMethods string) {
	request := http.Request{URL: &url.URL{Scheme: "http", Host: "127.0.0.1:8687", Path: requestPath},
		Method: "OPTIONS"}
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, &request)
	assert.Equal(t, allowedOrigin, writer.Header().Get(AccessControlAllowOriginHeader))
	assert.Equal(t, allowedHeader, writer.Header().Get(AccessControlAllowHeadersHeader))
	expectedMethodsList := strings.Split(allowedMethods, ",")
	actualMethodsList := strings.Split(writer.Header().Get(AccessControlAllowMethodsHeader), ",")
	testingutils.CheckStrings(t, expectedMethodsList, actualMethodsList, false, true)
}

func checkGinRouteCors(t *testing.T, router *gin.Engine, method string, requestPath string, allowedOrigin string) {
	request := http.Request{URL: &url.URL{Scheme: "http", Host: "127.0.0.1:8687", Path: requestPath},
		Method: method}
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, &request)
	assert.Equal(t, allowedOrigin, writer.Header().Get(AccessControlAllowOriginHeader))
}
