package main

import (
	"fmt"
	"net/http"
	"shorten/internal/auth"
)

const serverPortNumber = "8081"
const serverPort = ":" + serverPortNumber

func main() {
	// conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr:    serverPort,
		Handler: router,
	}

	fmt.Println("Server start and listening port:", serverPortNumber)
	server.ListenAndServe()
}
