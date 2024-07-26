package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	endpointList := []string{
		"https://google.com",
		"https://fb.com",
		"https://instagram.com",
		"https://github.com",
	}

	for _, web := range endpointList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error while calling endpoint %s", endpoint)
	} else {
		fmt.Printf("Status code %v while calling endpoint %s\n", res.StatusCode, endpoint)
	}

}
