package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
var baseURL string = "https://www.saramin.co.kr/zf_user/search?search_area=main&search_done=y&search_optional_item=n&searchType=search&searchword=java"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 1; i < totalPages; i++ {
		pageUrl := baseURL + "&recruitPage=" + strconv.Itoa(i)
		fmt.Println("Requesting ", pageUrl)
	}
}

func getPages() int {
	pages := 0

	response, err := http.Get(baseURL)

	checkErr(err)
	checkCode(response)

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
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
