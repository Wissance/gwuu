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
	pathVariableSign     = ":"
)

func getRouteBasePath(path string) string {
	trimmedPath := path
	if path[len(path)-1] == '/' {
		trimmedPath = path[0 : len(path)-2]
	}

	basePath := trimmedPath
	basePathEndIndex := strings.LastIndex(trimmedPath, "/")
	// pathVariableSignIndex := strings.LastIndex(trimmedPath, pathVariableSign)
	if basePathEndIndex > 0 {
		// if pathVariableSignIndex > basePathEndIndex
		basePath = trimmedPath[0 : basePathEndIndex-1]
	}
	return basePath
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
