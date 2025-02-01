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
}

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
