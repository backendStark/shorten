package main

import (
	"fmt"
	"net/http"
)

const serverPortNumber = "8081"
const serverPort = ":" + serverPortNumber

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hello", helloHandler)

	server := http.Server{
		Addr:    serverPort,
		Handler: router,
	}

	fmt.Println("Server start and listening port:", serverPortNumber)
	server.ListenAndServe()
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}
