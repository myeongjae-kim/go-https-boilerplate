package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hrzon/go-vue-boilerplate/internal/app/webserver"
)

func main() {
	fmt.Println("Hello, world!")

	http.HandleFunc("/", webserver.RootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
