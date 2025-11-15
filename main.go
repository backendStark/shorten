package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const googleURL = "https://google.com/"

func main() {
	t := time.Now()
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		go doReqToGoogle(i, &wg)
	}

	wg.Wait()
	fmt.Println(time.Since(t))
}

func doReqToGoogle(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(googleURL)

	if err != nil {
		fmt.Println("error get response from google")
		return
	}

	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		fmt.Printf("%d. Response from google successful, status code: %d\n", index+1, res.StatusCode)
	} else {
		fmt.Printf("%d. error get response from google, status code: %d\n", index+1, res.StatusCode)
	}
}
