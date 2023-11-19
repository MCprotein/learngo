package main

import (
	"errors"
	"fmt"
	"net/http"

	trackingtime "github.com/MCprotein/learngo/url-checker/tracking-time"
)

type resultRequest struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	defer trackingtime.Trackingtime()()

	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://google.com",
		"https://www.soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	results := map[string]string{}
	channel := make(chan resultRequest)

	for _, url := range urls {
		go hitURL(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

}

/*
chan<-: send only channel
*/
func hitURL(url string, channel chan<- resultRequest) {
	response, err := http.Get(url)
	status := "OK"
	if err != nil || response.StatusCode >= 400 {
		status = "FAILED"
	}
	channel <- resultRequest{url: url, status: status}

}
