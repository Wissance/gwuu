package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGinHandleFuncWithCorsWithAnyOrigin(t *testing.T) {
	handler := NewGinBasedWebApiHandler(true, AnyOrigin)

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

	checkGinOptionRouteCors(t, handler.Router, realmResource+"/", AnyOrigin, "*", "OPTIONS,GET")
	checkGinOptionRouteCors(t, handler.Router, userResource+"/", AnyOrigin, "*", "OPTIONS,GET,POST")
}

func checkGinOptionRouteCors(t *testing.T, router *gin.Engine, requestPath string, allowedOrigin string,
	allowedHeader string, allowedMethods string) {
	request := http.Request{URL: &url.URL{Scheme: "http", Host: "127.0.0.1:8687", Path: requestPath},
		Method: "OPTIONS"}
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, &request)
	assert.Equal(t, allowedOrigin, writer.Header().Get(AccessControlAllowOriginHeader))
	assert.Equal(t, allowedHeader, writer.Header().Get(AccessControlAllowHeadersHeader))
	assert.Equal(t, allowedMethods, writer.Header().Get(AccessControlAllowMethodsHeader))
}
