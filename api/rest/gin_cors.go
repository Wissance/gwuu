package rest

import (
	g "github.com/gin-gonic/gin"
	"net/http"
)

type GinBasedWebApiHandler struct {
	corsConfig map[string][]string
	AllowCors  bool
	Origin     string
	Router     *g.Engine
}

func NewGinBasedWebApiHandler(allowCors bool, origin string, opts ...g.OptionFunc) *GinBasedWebApiHandler {
	handler := GinBasedWebApiHandler{
		AllowCors:  allowCors,
		Origin:     origin,
		corsConfig: map[string][]string{},
	}

	handler.Router = g.Default(opts...)

	return &handler
}

func (handler *GinBasedWebApiHandler) GET(path string, f func(ctx *g.Context)) {
	handler.Router.GET(path, f)
	handler.addCorsHandler(path, http.MethodGet)
}

func (handler *GinBasedWebApiHandler) POST(path string, f func(ctx *g.Context)) {
	handler.Router.POST(path, handler.handleWithCors(f))
	handler.addCorsHandler(path, http.MethodPost)
}

func (handler *GinBasedWebApiHandler) PUT(path string, f func(ctx *g.Context)) {
	handler.Router.PUT(path, handler.handleWithCors(f))
	handler.addCorsHandler(path, http.MethodPut)
}

func (handler *GinBasedWebApiHandler) PATCH(path string, f func(ctx *g.Context)) {
	handler.Router.PATCH(path, handler.handleWithCors(f))
	handler.addCorsHandler(path, http.MethodPatch)
}

func (handler *GinBasedWebApiHandler) DELETE(path string, f func(ctx *g.Context)) {
	handler.Router.DELETE(path, handler.handleWithCors(f))
	handler.addCorsHandler(path, http.MethodDelete)
}

func (handler *GinBasedWebApiHandler) addCorsHandler(path string, method string) {
	if handler.AllowCors {
		corsMethods := make([]string, 0)
		routesInfo := handler.Router.Routes()
		// 1. Get base path /api/zone/realm and /api/zone/realm/1 has base -> /api/zone/
		basePath := getRouteBasePath(path)
		for _, r := range routesInfo {
			// 2. Collect all methods with same base path i.e. POST /api/zone/realm + 2 get methods /api/zone/realm and /api/zone/realm/1
			routeBasePath := getRouteBasePath(r.Path)
			if basePath == routeBasePath {
				corsMethods = append(corsMethods, r.Method)
			}
		}
		// 3. Check Cors Handler Configured
		optionHandlerExists := false
		for route, _ := range handler.corsConfig {
			if route == basePath {
				optionHandlerExists = true
			}
		}
		// 4. Create OPTION Handler if it does not exist
		if !optionHandlerExists {
			corsMethods = append(corsMethods, http.MethodOptions)
			// 4.1 Add Preflight Request Handler
		}
		// 5. Modify OPTION Methods Verbs list
		handler.corsConfig[basePath] = corsMethods
	}
}

// handleWithCors function that adds CORS headers before handler func is called
func (handler *GinBasedWebApiHandler) handleWithCors(f func(ctx *g.Context)) g.HandlerFunc {
	return func(ctx *g.Context) {
		if handler.AllowCors {
			handler.setCorsHeaders(ctx, handler.Origin)
		}

		f(ctx)
	}
}

func (handler *GinBasedWebApiHandler) setCorsHeaders(ctx *g.Context, origin string) {
	ctx.Writer.Header().Set(AccessControlAllowHeadersHeader, AllowAllHeaderValues)
	ctx.Writer.Header().Set(AccessControlAllowOriginHeader, origin)
}
