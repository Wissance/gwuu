package rest

import "strings"

const (
	// AccessControlAllowMethodsHeader - CORS Header that says what HTTP Methods are allowed to specific endpoint
	AccessControlAllowMethodsHeader = "Access-Control-Allow-Methods"
	AccessControlAllowOriginHeader  = "Access-Control-Allow-Origin"
	AccessControlAllowHeadersHeader = "Access-Control-Allow-Headers"
	// Value that allow all headers
	AnyOrigin            = "*"
	AllowAllHeaderValues = "*"
	optionsRouteSuffix   = "opts"
)

func getRouteBasePath(path string) string {
	trimmedPath := path
	if path[len(path)-1] == '/' {
		trimmedPath = path[0 : len(path)-2]
	}

	basePath := trimmedPath
	basePathEndIndex := strings.LastIndex(trimmedPath, "/")
	if basePathEndIndex > 0 {
		basePath = trimmedPath[0 : basePathEndIndex-1]
	}
	return basePath
}
