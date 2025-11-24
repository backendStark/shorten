package main

import (
	"fmt"
	"net/http"
)

const serverPortNumber = "8081"
const serverPort = ":" + serverPortNumber

func main() {
	fmt.Println("Server start and listening port:", serverPortNumber)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(serverPort, nil)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}
