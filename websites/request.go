package websites

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func request(url string, client *http.Client) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "ReaperJobsScraper v1.0 https://www.github.com/reaper47/jobs-scraper - This bot scrapes some job websites periodically for opporunities.")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
