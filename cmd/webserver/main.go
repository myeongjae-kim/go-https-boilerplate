package main

import (
	"log"
	"net/http"

	"github.com/hrzon/go-vue-boilerplate/internal/app/webserver"
	"github.com/hrzon/go-vue-boilerplate/pkg/logger"
)

func main() {
	logFile, err := logger.InitLoggerWithLogFileName("log")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer logFile.Close()

	http.HandleFunc("/", webserver.RootHandler)

	//TODO: Give an option to run a server with port number 8080 or 80
	log.Fatal(http.ListenAndServe(":8080", nil))
}
