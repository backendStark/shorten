package main

import (
	"fmt"
	"net/http"
	"time"
)

const googleURL = "https://google.com/"

func main() {
	t := time.Now()
	code := make(chan int)
	go doReqToGoogle(code)
	<-code
	fmt.Println("Program has been worked", time.Since(t), "seconds")
}

func doReqToGoogle(code chan int) {
	res, err := http.Get(googleURL)

	if err != nil {
		fmt.Println("error get response from google")
		return
	}

	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		fmt.Println("succesful request, get status code:", res.StatusCode)
		code <- res.StatusCode
		return
	} else {
		return
	}
}
