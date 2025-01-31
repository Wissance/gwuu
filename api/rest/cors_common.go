package rest

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
