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

func (handler *GinBasedWebApiHandler) GET(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.GET(path, handler.handleWithCors(f))
	} else {
		routerGroup.GET(path, handler.handleWithCors(f))
	}

	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path), http.MethodGet)
}

func (handler *GinBasedWebApiHandler) POST(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.POST(path, handler.handleWithCors(f))
	} else {
		routerGroup.POST(path, handler.handleWithCors(f))
	}
	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path), http.MethodPost)
}

func (handler *GinBasedWebApiHandler) PUT(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.PUT(path, handler.handleWithCors(f))
	} else {
		routerGroup.PUT(path, handler.handleWithCors(f))
	}
	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path), http.MethodPut)
}

func (handler *GinBasedWebApiHandler) PATCH(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.PATCH(path, handler.handleWithCors(f))
	} else {
		routerGroup.PATCH(path, handler.handleWithCors(f))
	}
	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path), http.MethodPatch)
}

func (handler *GinBasedWebApiHandler) DELETE(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.DELETE(path, handler.handleWithCors(f))
	} else {
		routerGroup.DELETE(path, handler.handleWithCors(f))
	}
	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path), http.MethodDelete)
}

func (handler *GinBasedWebApiHandler) addCorsHandler(path string, method string) {
	if handler.AllowCors {
		corsMethods := make([]string, 0)
		routesInfo := handler.Router.Routes()
		// 1. Find route by appropriate path
		for _, r := range routesInfo {
			// 2. Collect all methods with same path i.e. POST /api/zone/realm + 2 get methods /api/zone/realm and /api/zone/realm/1
			if path == r.Path {
				corsMethods = append(corsMethods, r.Method)
			}
		}
		// 3. Check Cors Handler Configured
		optionHandlerExists := false
		for route, _ := range handler.corsConfig {
			if route == path {
				optionHandlerExists = true
				break
			}
		}
		// 4. Create OPTION Handler if it does not exist
		if !optionHandlerExists {
			// 4.1 Add Preflight Request Handler
			corsMethods = append([]string{http.MethodOptions}, corsMethods...)
			handler.Router.OPTIONS(path, handler.handlePreflightReq)
		}
		// 5. Modify OPTION Methods Verbs list
		handler.corsConfig[path] = corsMethods
	}
}

func (handler *GinBasedWebApiHandler) handlePreflightReq(ctx *g.Context) {
	route := handler.getRouteInfo(ctx, http.MethodOptions)
	if route != nil {
		methods := handler.corsConfig[route.Path]
		methodsStr := join(methods, ",")
		handler.enableCors(ctx, handler.Origin, methodsStr)
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

func (handler *GinBasedWebApiHandler) enableCors(ctx *g.Context, origin string, methods string) {
	handler.setCorsHeaders(ctx, origin)
	ctx.Writer.Header().Set(AccessControlAllowMethodsHeader, methods)
}

func (handler *GinBasedWebApiHandler) getRouteInfo(ctx *g.Context, method string) *g.RouteInfo {
	// handlerName := ctx.HandlerName()
	path := ctx.FullPath()
	routesInfo := handler.Router.Routes()
	for _, r := range routesInfo {
		if r.Path == path && r.Method == method {
			return &r
		}
	}
	return nil
}

func (handler *GinBasedWebApiHandler) getRouteBasePath(routerGroup *g.RouterGroup, path string) string {
	if routerGroup == nil {
		return path
	}
	return routerGroup.BasePath() + path
}
