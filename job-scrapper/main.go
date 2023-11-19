package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
var baseURL string = "https://www.saramin.co.kr/zf_user/search?search_area=main&search_done=y&search_optional_item=n&searchType=search&searchword=golang"

func main() {
	getPages()
}

func getPages() int {
	response, err := http.Get(baseURL)

	checkErr(err)
	checkCode(response)

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)

	doc.Find("pagination").Each()

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln("Http Error", err)
	}
}

func checkCode(response *http.Response) {
	if response.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", response.StatusCode)
	}
}
