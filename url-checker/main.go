package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
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
	/*
		make: map을 만들어주는 함수
	*/
	// results2 := make(map[string]string)

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)

		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	response, err := http.Get(url)

	if err != nil || response.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
