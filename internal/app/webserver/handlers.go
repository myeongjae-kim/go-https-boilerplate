package webserver

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	root = "./web/dist/"
)

// RootHandler is an handler to send static files
func RootHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/"):]

	source, err := ioutil.ReadFile(root + path)
	if err != nil {
		source, err = ioutil.ReadFile(root + path + "/index.html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		path += "index.html"
	}

	// Set content type to css if required file's extension is css.
	if len(path) >= 4 && path[len(path)-4:] == ".css" {
		w.Header().Set("Content-Type", "text/css")
	}

	w.Write(source)

	//TODO: Log more detailed information.
	log.Println("(rootHandler) The requested file has been sent: ", root+path)
}
