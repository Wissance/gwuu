package rest

import (
	m "github.com/gorilla/mux"
	"github.com/wissance/stringFormatter"
	"net/http"
)

// MuxBasedWebApiHandler is a struct that encapsulates Router and CORS settings
type MuxBasedWebApiHandler struct {
	corsConfig map[string][]string
	AllowCors  bool
	Origin     string
	Router     *m.Router
}

// NewMuxBasedWebApiHandler
/* This function creates instance of MuxBasedWebApiHandler and inits properties with arguments values
 * Parameters:
 *    - allowCors - represents should be course configured (true) or not
 *    - origin - represents ip/domain or * (AnyOrigin) name which will be used response headers
 */
func NewMuxBasedWebApiHandler(allowCors bool, origin string) *MuxBasedWebApiHandler {
	handler := &MuxBasedWebApiHandler{
		AllowCors:  allowCors,
		Origin:     origin,
		Router:     m.NewRouter(),
		corsConfig: map[string][]string{},
	}

	if allowCors {
		handler.Router.Use(m.CORSMethodMiddleware(handler.Router))
	}

	return handler
}

// EnableCorsAllOrigin
/* This function sets CORS headers to allow any origin. In the future it should take array with allowed origins (ip addresses and/or hostnames)
 * Parameters:
 *     - respWriter - gorilla/mux response writer
 * Returns nothing
 */
func EnableCorsAllOrigin(respWriter *http.ResponseWriter) {
	EnableCors(respWriter, AllowAllHeaderValues, AllowAllHeaderValues)
}

// EnableCors
/* This function sets CORS headers to specified origin. In the future it should take array with allowed origins (ip addresses and/or hostnames)
 * Parameters:
 *     - respWriter - gorilla/mux response writer
 *     - origin - domain i.e. example.com or http://127.0.0.1:3000
 *     - methods - allowed methods i.e. GET, POST, OPTIONS
 * Returns nothing
 */
func EnableCors(respWriter *http.ResponseWriter, origin string, methods string) {
	addCorsHeaders(respWriter, origin)
	(*respWriter).Header().Set(AccessControlAllowMethodsHeader, methods)
}

func addCorsHeaders(respWriter *http.ResponseWriter, origin string) {
	(*respWriter).Header().Set(AccessControlAllowHeadersHeader, AllowAllHeaderValues)
	(*respWriter).Header().Set(AccessControlAllowOriginHeader, origin)
}

func (handler *MuxBasedWebApiHandler) handlePreflightReq(respWriter http.ResponseWriter, request *http.Request) {
	route := m.CurrentRoute(request)
	if route != nil {
		optionRouteName := route.GetName()
		// stringFormatter.Format("{0}_{1}", request.URL.Path, optionsRouteSuffix)
		methods := handler.corsConfig[optionRouteName]
		methodsStr := join(methods, ",")
		EnableCors(&respWriter, handler.Origin, methodsStr)
	}
}

// handleWithCors function that adds CORS headers before handler func is called
func (handler *MuxBasedWebApiHandler) handleWithCors(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if handler.AllowCors {
			addCorsHeaders(&writer, handler.Origin)
		}
		f(writer, request)
	}
}

// HandleFunc
/* This is a Proxy function that assign handler to handle specific route by url but also simultaneously it configures CORS handler.
 * This function is almost equal to mux.Router.HandleFunc except fact that we passing
 * We are working here with route names for OPTIONS handler i.e. we have REST resource /api/good and 2 separate handlers for GET and POST
 * therefore for proper CORS handle we should respond on OPTIONS /api/good with empty body and header AccessControlAllowMethodsHeader with
 * values OPTIONS, GET, POST. Our HandleFunc allow to reduce a complexity of router config because using our HandleFunc we take service on
 * handling OPTIONS method by our HandleFunc.
 * Parameters:
 *     - router - router to which we assign handler func this is implemented for sub routers supports
 *     - path - url of route (request)
 *     - f - handler function that handles request
 * Return created route like router.HandleFunc do
 */
func (handler *MuxBasedWebApiHandler) HandleFunc(router *m.Router, path string, f func(http.ResponseWriter, *http.Request), handlerMethods ...string) *m.Route {
	// 1. Create Route ...
	route := router.HandleFunc(path, handler.handleWithCors(f)).Methods(handlerMethods...)
	actualRoutePath := path
	// This code taking into account that router is a SubRouter
	s, err := route.GetPathTemplate()
	if err == nil {
		actualRoutePath = s
	}
	if handler.AllowCors {
		// 2. Create Options route
		optionRouteName := stringFormatter.Format("{0}_{1}", actualRoutePath, optionsRouteSuffix)
		optionsRoute := router.GetRoute(optionRouteName)
		if optionsRoute == nil {
			// there is no Route with such name, so we could easily create it and assign methods = "OPTIONS" + handlerMethods
			handler.corsConfig[optionRouteName] = []string{"OPTIONS"}
			// assign OPTIONS Handler
			router.HandleFunc(path, handler.handlePreflightReq).Methods("OPTIONS").Name(optionRouteName)
		}
		// combine with handlerMethods
		handler.corsConfig[optionRouteName] = append(handler.corsConfig[optionRouteName], handlerMethods...)
	}
	return route
}
