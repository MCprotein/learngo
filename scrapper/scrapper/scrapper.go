package scrapper

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

func Scrape(term string) {
	baseURL := "https://www.saramin.co.kr/zf_user/search?search_area=main&search_done=y&search_optional_item=n&searchType=search&searchword=" + term
	jobs := []extractedJob{}
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	for i := 1; i < totalPages; i++ {
		go getPage(i, baseURL, c)

	}

	for i := 1; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	csvChannel := make(chan error)
	go writeJobs(jobs, csvChannel)

	for i := 0; i < len(jobs); i++ {
		jwErr := <-csvChannel
		checkErr(jwErr)
	}

	fmt.Println("Done, extracted ", len(jobs))
}

func getPage(page int, url string, mainC chan<- []extractedJob) {
	jobs := []extractedJob{}
	c := make(chan extractedJob)

	pageUrl := url + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting ", pageUrl)
	response, err := http.Get(pageUrl)
	checkErr(err)
	checkCode(response)

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func getPages(url string) int {
	pages := 0

	response, err := http.Get(url)

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

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")

	title := CleanString(card.Find(".area_job .job_tit>a").Text())

	location := CleanString(card.Find(".job_condition>span>a").Text())

	company := CleanString(card.Find(".area_corp .corp_name>a").Text())

	career := ""
	education := ""
	salary := ""
	card.Find(".job_condition>span").Each(func(i int, description *goquery.Selection) {
		if i == 0 {
			return
		}

		if i == 1 {
			career = CleanString(description.Text())
		}

		switch i {
		case 1:
			career = CleanString(description.Text())
		case 2:
			education = CleanString(description.Text())
		case 4:
			salary = CleanString(description.Text())
		}

	})

	c <- extractedJob{id: id, title: strings.TrimSpace(title), location: location, salary: salary, career: career, company: company, education: education}

}

func writeJobs(jobs []extractedJob, csvChannel chan<- error) {
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
		csvChannel <- jwErr
	}

}
