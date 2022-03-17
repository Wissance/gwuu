package rest

import (
	m "github.com/gorilla/mux"
	"github.com/wissance/stringFormatter"
	"net/http"
)

const (
	// AccessControlAllowMethodsHeader - CORS Header that says what HTTP Methods are allowed to specific endpoint
	AccessControlAllowMethodsHeader = "Access-Control-Allow-Methods"
	AccessControlAllowOriginHeader = "Access-Control-Allow-Origin"
	AccessControlAllowHeadersHeader = "Access-Control-Allow-Headers"
	// Value that allow all headers
	AnyOrigin = "*"
	AllowAllHeaderValues = "*"
	optionsRouteSuffix = "opts"
)

type WebApiHandler struct {
	corsConfig map[string][]string
	AllowCors bool
	Origin string
	Router *m.Router
}

// NewWebApiHandler
/* This function creates instance of WebApiHandler and inits properties with arguments values
 * Parameters:
 *    - allowCors - represents should be course configured (true) or not
 *    - origin - represents ip/domain or * (AnyOrigin) name which will be used response headers
 */
func NewWebApiHandler(allowCors bool, origin string) *WebApiHandler {
	handler := &WebApiHandler{
		AllowCors: allowCors,
		Origin: origin,
		Router: m.NewRouter(),
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
	(*respWriter).Header().Set(AccessControlAllowHeadersHeader, AllowAllHeaderValues)
	(*respWriter).Header().Set(AccessControlAllowOriginHeader, origin)
	(*respWriter).Header().Set(AccessControlAllowMethodsHeader, methods)
}

// HandleFunc
/* This is a Proxy function that assign handler to handle specific route by url but also simultaneously it configures CORS handler.
 * We are working here with route names for OPTIONS handler i.e. we have REST resource /api/good and 2 separate handlers for GET and POST
 * therefore for proper CORS handle we should respond on OPTIONS /api/good with empty body and header AccessControlAllowMethodsHeader with
 * values OPTIONS, GET, POST. Our HandleFunc allow to reduce a complexity of router config because using our HandleFunc we take service on
 * handling OPTIONS method by our HandleFunc.
 * Parameters:
 *     - router - router to which we assign handler func this is implemented for sub routers supports
 *     - path - url
 *     - f - handler function that handles request
 * Return *Route
 */
func (handler *WebApiHandler) HandleFunc(router *m.Router, path string, f func(http.ResponseWriter, *http.Request), handlerMethods ...string) *m.Route {
	// 1. Create Route ...
	route := router.HandleFunc(path, f).Methods(handlerMethods...)
	if handler.AllowCors {
		// 2. Create Options route
		optionRouteName := stringFormatter.Format("{0}_{1}", path, optionsRouteSuffix)
		optionsRoute := router.GetRoute(optionRouteName)
		if optionsRoute == nil {
			// there is no Route with such name, so we could easily create it and assign methods = "OPTIONS" + handlerMethods
			handler.corsConfig[optionRouteName] = []string{"OPTIONS"}
			// assign OPTIONS Handler
			router.HandleFunc(path, handler.handleCors).Methods("OPTIONS").Name(optionRouteName)
		}
		// combine with handlerMethods
		handler.corsConfig[optionRouteName] = append(handler.corsConfig[optionRouteName], handlerMethods...)
	}
	return route
}

func (handler *WebApiHandler) handleCors(respWriter http.ResponseWriter, request *http.Request) {
	optionRouteName := stringFormatter.Format("{0}_{1}", request.URL.Path, optionsRouteSuffix)
	methods := handler.corsConfig[optionRouteName]
	methodsStr := join(methods, ",")
	EnableCors(&respWriter, handler.Origin, methodsStr)
}

func join(values []string, separator string) string {
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + separator
		}
	}
	return line
}
