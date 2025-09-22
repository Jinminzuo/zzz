package main

import (
	"log"
	"net/http"
	"wxcloudrun-golang/service"
)

func main() {

	http.HandleFunc("/", service.HelloWorldHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
