package main

import (
	"fmt"
	"net/http"
	"time"
)

const googleURL = "https://google.com/"

func main() {
	t := time.Now()
	for i := range 10 {
		res, err := http.Get(googleURL)

		if err != nil {
			fmt.Println("error get response from google")
			return
		}

		defer res.Body.Close()

		if res.StatusCode >= 200 && res.StatusCode < 300 {
			fmt.Printf("%d. Response from google successful, status code: %d\n", i+1, res.StatusCode)
		} else {
			fmt.Printf("%d. error get response from google, status code: %d\n", i+1, res.StatusCode)
		}
	}
	fmt.Println(time.Since(t))
}
