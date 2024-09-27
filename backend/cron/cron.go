package cron

import (
	"EQA/backend/app/service"

	"github.com/robfig/cron/v3"
)

func NewCron(sv service.ProgramService) *cron.Cron {
	c := cron.New()

	// Daily 8 AM
	c.AddFunc("0 8 * * *", func() {
		sv.Reminder()
	})

	return c
}
