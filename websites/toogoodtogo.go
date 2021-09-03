package websites

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/reaper47/jobs-scraper/model"
)

const BASE_URL = "https://toogoodtogo.org"

func ScrapeTooGoodToGo(client *http.Client) (*model.JobsMetaData, error) {
	doc, err := request(BASE_URL+"/en/careers/opportunities?location=Canada", client)
	if err != nil {
		return nil, err
	}

	var jobs []*model.Job
	doc.Find("a").Each(func(i1 int, a *goquery.Selection) {
		href, exists := a.Attr("href")
		if exists {
			a.Find(".position-link").Each(func(i2 int, span *goquery.Selection) {
				jobs = append(jobs, &model.Job{
					Position: span.Text(),
					URL:      BASE_URL + href,
				})
			})
		}
	})

	return &model.JobsMetaData{
		Title: "Too Good To Go",
		Logo:  "https://tgtg-mkt-cms-prod.s3.eu-west-1.amazonaws.com/6610/TGTG_Icon_2000x2000_RGB.png",
		Jobs:  jobs,
	}, nil
}
