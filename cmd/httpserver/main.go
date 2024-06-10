package main

import (
	"45HW/pkg/config"
	"45HW/pkg/http_server"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	router := http_server.NewRouter()
	log.Fatal(http.ListenAndServe(":"+config.Conf.HTTPServerPort, router))
}
