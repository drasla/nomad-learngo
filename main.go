package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	c := make(chan result)
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
		go hitURL(url, c)
	}
}

func hitURL(url string, c chan<- result) {
	fmt.Println("Checking: ", url)
	response, err := http.Get(url)
	if err != nil || response.StatusCode >= 400 {
		c <- result{url: url, status: "FAILDED"}
	} else {
		c <- result{url: url, status: "OK"}
	}
}
