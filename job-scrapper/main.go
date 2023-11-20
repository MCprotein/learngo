package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id        string
	location  string
	title     string
	salary    string
	career    string
	company   string
	education string
}

// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
var baseURL string = "https://www.saramin.co.kr/zf_user/search?search_area=main&search_done=y&search_optional_item=n&searchType=search&searchword=java"

func main() {
	jobs := []extractedJob{}
	totalPages := getPages()

	for i := 1; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)

	}
	writeJobs(jobs)
	fmt.Println("Done, extracted ", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Location", "Title", "Salary", "Career", "Company", "Education"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?isMypage=no&rec_idx=" + job.id, job.location, job.title, job.salary, job.career, job.company, job.education}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

}

func getPage(page int) []extractedJob {
	jobs := []extractedJob{}
	pageUrl := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting ", pageUrl)
	response, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(response)

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
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

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")

	title := cleanString(card.Find(".area_job .job_tit>a").Text())

	location := cleanString(card.Find(".job_condition>span>a").Text())

	company := cleanString(card.Find(".area_corp .corp_name>a").Text())

	career := ""
	education := ""
	salary := ""
	card.Find(".job_condition>span").Each(func(i int, description *goquery.Selection) {
		if i == 0 {
			return
		}

		if i == 1 {
			career = cleanString(description.Text())
		}

		switch i {
		case 1:
			career = cleanString(description.Text())
		case 2:
			education = cleanString(description.Text())
		case 4:
			salary = cleanString(description.Text())
		}

	})

	return extractedJob{id: id, title: strings.TrimSpace(title), location: location, salary: salary, career: career, company: company, education: education}

}
