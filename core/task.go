package core

import (
	"context"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

// IsFriday checks whether it is the following Friday after last Friday.
func IsFriday() bool {
	now := time.Now()
	day := now.Day()

	data, err := ioutil.ReadFile("schedule.data")
	if err != nil {
		content := fmt.Sprint(day)
		ioutil.WriteFile("schedule.data", []byte(content), 0777)
		return false
	}
	prevDay, _ := strconv.Atoi(string(data))

	if now.Weekday() == time.Friday && day > prevDay {
		content := fmt.Sprint(day)
		ioutil.WriteFile("schedule.data", []byte(content), 0777)
		return true
	}
	return false
}

// Cron makes a cron job.
func Cron(ctx context.Context, startTime time.Time, delay time.Duration) <-chan time.Time {
	stream := make(chan time.Time, 1)

	if !startTime.IsZero() {
		diff := time.Until(startTime)
		if diff < 0 {
			total := diff - delay
			times := total / delay * -1

			startTime = startTime.Add(times * delay)
		}
	}

	go func() {
		t := <-time.After(time.Until(startTime))
		stream <- t

		ticker := time.NewTicker(delay)
		defer ticker.Stop()

		for {
			select {
			case t2 := <-ticker.C:
				stream <- t2
			case <-ctx.Done():
				close(stream)
				return
			}
		}
	}()

	return stream
}
