package core

import (
	"log"
	"net/http"
	"time"

	"github.com/reaper47/jobs-scraper/config"
	"github.com/reaper47/jobs-scraper/email"
	"github.com/reaper47/jobs-scraper/model"
	"github.com/reaper47/jobs-scraper/websites"
)

func SendMail() {
	subject := "Anya's Weekly Jobs Digest"
	r := email.NewRequest(config.Config.To, subject)
	if err := r.Send("template/jobs.html", getJobs()); err != nil {
		log.Fatal(err)
	}
	log.Println("Email sent")
}

func getJobs() *model.Websites {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	jobsTooGoodToGo, err := websites.ScrapeTooGoodToGo(client)
	if err != nil {
		log.Fatal(err)
	}

	jobs := []*model.JobsMetaData{jobsTooGoodToGo}
	return &model.Websites{Metadata: jobs}
}
