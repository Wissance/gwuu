package rest

import (
	m "github.com/gorilla/mux"
	"net/http"
)

const (
	AccessControlAllowMethodsHeader = "Access-Control-Allow-Methods"
	AccessControlAllowOriginHeader = "Access-Control-Allow-Origin"
	AccessControlAllowHeadersHeader = "Access-Control-Allow-Headers"
	AllowAllHeaderValues = "*"
)

/*type WebApi interface {
	//Create()
}*/

type WebApiHandler struct {
	Origin string
	Router *m.Router
}

// EnableCorsAllOrigin
/* This function sets CORS headers to allow any origin. In the future it should take array with allowed origins (ip addresses and/or hostnames)
 * Parameters:
 *     - respWriter - gorilla/mux response writer
 * Returns nothing
 */
func EnableCorsAllOrigin(respWriter *http.ResponseWriter) {
	EnableCors(respWriter, AllowAllHeaderValues)
}

// EnableCors
/* This function sets CORS headers to specified origin. In the future it should take array with allowed origins (ip addresses and/or hostnames)
 * Parameters:
 *     - respWriter - gorilla/mux response writer
 *     - origin - domain i.e. example.com or http://127.0.0.1:3000
 * Returns nothing
 */
func EnableCors(respWriter *http.ResponseWriter, origin string) {
	(*respWriter).Header().Set(AccessControlAllowOriginHeader, origin)
	(*respWriter).Header().Set(AccessControlAllowHeadersHeader, AllowAllHeaderValues)
}

func (handler *WebApiHandler) HandleFunc(url string, f func(http.ResponseWriter, *http.Request), allowCors bool) {
	// 1. Create Route ...
	// 2. Create Options route

}
