package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/myeongjae-kim/go-https-boilerplate/pkg/webserver/handlers/headers"
)

// The server have to be run in root directory of a project.
// The default root directory is './web'. See https://github.com/golang-standards/project-layout
var root = "./website/"

// SetRootDirectory sets a default direcrtory of the RootHandler to search resources.
func SetRootDirectory(rootDirectory string) {
	root = rootDirectory
	if root[len(root)-1] != '/' {
		root += "/"
	}
	log.Println("Set root directory to : '" + root + "'")
}

// RootHandler is an handler to send static files
func RootHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/"):]

	source, err := ioutil.ReadFile(root + filePath)
	if err != nil {
		source, err = ioutil.ReadFile(root + filePath + "/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		filePath += "index.html"
	}

	// Set response headers
	headers.SetDefaultHeaders(w)
	headers.SetContentTypeHeader(w, filePath)

	// Send the response
	w.Write(source)

	//TODO: Log more detailed information.
	log.Println("(rootHandler) The requested file has been sent: ", root+filePath)
}
