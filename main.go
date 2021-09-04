package main

import (
	"context"
	"time"

	"github.com/reaper47/jobs-scraper/config"
	"github.com/reaper47/jobs-scraper/core"
)

func main() {
	config.InitConfig()

	ctx := context.Background()
	for range core.Cron(ctx, time.Now(), 12*time.Hour) {
		if core.IsFriday() {
			core.SendMail()
		}
	}
}
