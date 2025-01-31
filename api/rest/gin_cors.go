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
		handler.Router.GET(path, f)
	} else {
		routerGroup.GET(path, f)
	}

	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path))
}

func (handler *GinBasedWebApiHandler) POST(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.POST(path, handler.handleWithCors(f))
	} else {
		routerGroup.POST(path, f)
	}
	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path))
}

func (handler *GinBasedWebApiHandler) PUT(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.PUT(path, handler.handleWithCors(f))
	} else {
		routerGroup.PUT(path, f)
	}
	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path))
}

func (handler *GinBasedWebApiHandler) PATCH(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.PATCH(path, handler.handleWithCors(f))
	} else {
		routerGroup.PATCH(path, f)
	}
	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path))
}

func (handler *GinBasedWebApiHandler) DELETE(routerGroup *g.RouterGroup, path string, f func(ctx *g.Context)) {
	if routerGroup == nil {
		handler.Router.DELETE(path, handler.handleWithCors(f))
	} else {
		routerGroup.DELETE(path, handler.handleWithCors(f))
	}
	handler.addCorsHandler(handler.getRouteBasePath(routerGroup, path))
}

func (handler *GinBasedWebApiHandler) addCorsHandler(path string) {
	if handler.AllowCors {
		corsMethods := make([]string, 0)
		routesInfo := handler.Router.Routes()
		// 1. Get base path /api/zone/realm and /api/zone/realm/1 has base -> /api/zone/
		// todo(UMV): ???
		basePath := getRouteBasePath(path)
		for _, r := range routesInfo {
			// 2. Collect all methods with same base path i.e. POST /api/zone/realm + 2 get methods /api/zone/realm and /api/zone/realm/1
			// todo(UMV): ???
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
				break
			}
		}
		// 4. Create OPTION Handler if it does not exist
		if !optionHandlerExists {
			corsMethods = append(corsMethods, http.MethodOptions)
			// 4.1 Add Preflight Request Handler
			handler.Router.OPTIONS(path, handler.handlePreflightReq)
		}
		// 5. Modify OPTION Methods Verbs list
		handler.corsConfig[basePath] = corsMethods
	}
}

func (handler *GinBasedWebApiHandler) handlePreflightReq(ctx *g.Context) {
	route := handler.getRouteInfo(ctx)
	//m.CurrentRoute(request)
	if route != nil {
		routeBasePath := getRouteBasePath(route.Path)
		// stringFormatter.Format("{0}_{1}", request.URL.Path, optionsRouteSuffix)
		methods := handler.corsConfig[routeBasePath]
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

func (handler *GinBasedWebApiHandler) getRouteInfo(ctx *g.Context) *g.RouteInfo {
	handlerName := ctx.HandlerName()
	routesInfo := handler.Router.Routes()
	for _, r := range routesInfo {
		if r.Handler == handlerName {
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
