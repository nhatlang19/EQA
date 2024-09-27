package cron

import (
	"EQA/backend/app/service"

	"github.com/robfig/cron/v3"
)

func NewCron(sv service.ProgramService) *cron.Cron {
	c := cron.New()

	// Add a cron job that runs every minute
	c.AddFunc("* * * * *", func() {
		sv.Reminder()
	})

	return c
}
