package main

import (
	"log"
	"net/http"
	"simple/internal/hellower"
	"simple/restapi"
)

func main() {
	hello := &hellower.Hellower{}

	h, err := restapi.Handler(restapi.Config{
		HelloAPI: hello,
	})
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", h))
}
