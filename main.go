package main

import (
	"context"
	"log"
	"time"

	"github.com/reaper47/jobs-scraper/config"
	"github.com/reaper47/jobs-scraper/core"
)

func main() {
	log.Println("Program started")
	config.InitConfig()

	ctx := context.Background()
	for range core.Cron(ctx, time.Now(), 12*time.Hour) {
		if core.IsFriday() {
			core.SendMail()
		}
	}
}
