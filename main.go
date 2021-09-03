package main

import (
	"log"
	"net/http"
	"time"

	"github.com/reaper47/jobs-scraper/config"
	"github.com/reaper47/jobs-scraper/email"
	"github.com/reaper47/jobs-scraper/model"
	"github.com/reaper47/jobs-scraper/websites"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	jobsTooGoodToGo, err := websites.ScrapeTooGoodToGo(client)
	if err != nil {
		log.Fatal(err)
	}

	jobs := []*model.JobsMetaData{jobsTooGoodToGo}
	websites := &model.Websites{Metadata: jobs}

	subject := "Anya's Weekly Jobs Digest"
	to := []string{config.GetEnvVar("to")}
	r := email.NewRequest(to, subject)
	if err := r.Send("template/jobs.html", websites); err != nil {
		log.Fatal(err)
	}

}
