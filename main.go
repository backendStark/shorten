package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	path := flag.String("file", "urls.txt", "path to URL file")
	flag.Parse()

	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}

	urlSlice := strings.Split(string(file), "\n")
	respCh := make(chan int)
	errCh := make(chan error)

	for _, url := range urlSlice {
		go ping(url, respCh, errCh)
	}

	for i := 0; i < len(urlSlice); i++ {
		res := <- respCh
		fmt.Println(res)
		errRes := <- errCh
		fmt.Println(errRes)
	}
}

func ping(url string, respCh chan int, errCh chan error) {
	resp, err := http.Get(url)

	if err != nil {
		errCh <- err
		return
	}

	respCh <- resp.StatusCode
}
