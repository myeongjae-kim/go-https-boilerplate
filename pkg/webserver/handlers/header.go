package handlers

import (
	"path"
	"strings"
)

var contentType map[string]string

func init() {
	contentType = make(map[string]string)
	contentType[".html"] = "text/html"
	contentType[".css"] = "text/css"

	contentType[".svg"] = "image/svg+xml"
}

func getContentTypeHeader(filePath string) string {
	return contentType[strings.ToLower(path.Ext(filePath))]
}
