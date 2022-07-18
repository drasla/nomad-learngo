package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instragram.com/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("Request failed")

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	response, err := http.Get(url)
	if err == nil || response.StatusCode >= 400 {
		fmt.Println(err, response.StatusCode)
		return errRequestFailed
	}
	return nil
}
