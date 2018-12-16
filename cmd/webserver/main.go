package main

import (
	"flag"
	"log"

	"github.com/hrzon/go-vue-boilerplate/pkg/logger"
	"github.com/hrzon/go-vue-boilerplate/pkg/webserver"
	"github.com/hrzon/go-vue-boilerplate/pkg/webserver/handlers"
)

var (
	flagRedirectToHTTPS = false
)

func parseFlags() {
	flag.BoolVar(
		&flagRedirectToHTTPS,
		"redirect-to-https",
		false,
		"if true, we redirect HTTP to HTTPS")

	flag.Parse()

	if flagRedirectToHTTPS {
		log.Println("flagRedirectToHTTPS is set.")
	}
}

func main() {
	logFile, err := logger.InitLoggerWithLogFileName("log")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer logFile.Close()

	parseFlags()

	// Set handlers
	handlerMap := make(webserver.HandlerMap)
	handlerMap["/"] = handlers.RootHandler

	// Set HTTPS hosts
	allowedHTTPSHosts := []string{
		"live.myeongjae.kim",
		"book.myeongjae.kim",
	}

	handlers.SetRootDirectory("./web/dist/")
	webserver.SetRedirectToHTTPS(flagRedirectToHTTPS)
	webserver.SetHTTPPort(8080)
	webserver.InitAndRunServers(handlerMap, allowedHTTPSHosts)
}
