package headers

import (
	"log"
	"net/http"
	"path"
	"strings"
)

var contentTypes map[string]string
var defaultHeaders map[string]string

func init() {
	// Set content types
	contentTypes = make(map[string]string)
	contentTypes[".html"] = "text/html"
	contentTypes[".css"] = "text/css"
	contentTypes[".svg"] = "image/svg+xml"
	contentTypes[".js"] = "application/javascript"

	// Set default headers
	defaultHeaders = make(map[string]string)
	// Set below header for all responses
	// https://blog.stackpath.com/accept-encoding-vary-important
	defaultHeaders["Vary"] = "Accept-Encoding"

	defaultHeaders["X-Frame-Options"] = "deny"
	defaultHeaders["X-Xss-Protection"] = "1; mode=block"
	defaultHeaders["X-Content-Type-Options"] = "nosniff"
}

// SetContentTypeHeader writes content type header to a response according to an extension of a file path
func SetContentTypeHeader(w *http.ResponseWriter, filePath string) {
	(*w).Header().Set(
		"Content-Type",
		contentTypes[strings.ToLower(path.Ext(filePath))],
	)
}

// SetDefaultHeaders write headers registered as defaults
func SetDefaultHeaders(w *http.ResponseWriter) {
	for k, v := range defaultHeaders {
		(*w).Header().Set(k, v)
	}
}

// AddDefaultHeader add headers to set as defaults
func AddDefaultHeader(headerName string, value string) {
	defaultHeaders[headerName] = value
}

// AddContentType add a content type to the contentTypes map
func AddContentType(extension string, value string) {
	if extension[0] != '.' {
		log.Fatalln("(AddContentType) The first character of an extension have to be '.', but:", extension)
	}

	contentTypes[extension] = value
}
