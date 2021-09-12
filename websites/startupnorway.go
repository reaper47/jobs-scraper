package websites

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/reaper47/jobs-scraper/model"
)

func ScrapeStartupNorway(client *http.Client) (*model.JobsMetaData, error) {
	doc, err := request(URL_STARTUPNORWAY, client)
	if err != nil {
		return nil, err
	}

	var jobs []*model.Job
	doc.Find("h2").Each(func(i int, h2Node *goquery.Selection) {
		position := h2Node.Text()
		h2Node.Parent().Parent().Next().Find("a").Each(func(i2 int, a *goquery.Selection) {
			href, exists := a.Attr("href")
			if exists && strings.Contains(href, "startupmatcher") {
				jobs = append(jobs, &model.Job{
					Position: position,
					URL:      href,
				})
			}
		})
	})

	return &model.JobsMetaData{
		Title: "Startup Norway",
		Logo:  "https://d33wubrfki0l68.cloudfront.net/f513ecfe45838b888e49751588a30349abaaf33c/579b3/static/c4fb56847ddc29789305af411d9ef570/14872/logo-positive.png",
		Jobs:  jobs,
	}, nil
}
