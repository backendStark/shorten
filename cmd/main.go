package main

import (
	"fmt"
	"net/http"
	"shorten/internal/hello"
)

const serverPortNumber = "8081"
const serverPort = ":" + serverPortNumber

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    serverPort,
		Handler: router,
	}

	fmt.Println("Server start and listening port:", serverPortNumber)
	server.ListenAndServe()
}
