package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/shohag121/LetMeKnow/github"
	"github.com/shohag121/LetMeKnow/notification"
)

func AddCronJob() {
	c := getJob()
	c.Start()
}

func RemoveCronJob() {
	c := getJob()
	c.Stop()
}

func getJob() *cron.Cron {
	c := cron.New()
	// every 5 minutes
	c.AddFunc("@every 0h5m0s", ProcessCronJob)
	return c
}

func ProcessCronJob() {
	list, err := github.GetUserNotifications()

	if err != nil {
		fmt.Println(err)
	}

	notification.Process(list)
}
