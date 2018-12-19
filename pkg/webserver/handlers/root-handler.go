package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

// The server have to be run in root directory of a project.
// The default root directory is './web'. See https://github.com/golang-standards/project-layout
var root = "./web/"

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

	// Set header according to its file extension.
	switch path.Ext(filePath) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Header().Set("Vary", "Accept-Encoding")
	}

	w.Write(source)

	//TODO: Log more detailed information.
	log.Println("(rootHandler) The requested file has been sent: ", root+filePath)
}
