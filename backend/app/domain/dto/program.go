package dto

import (
	model "EQA/backend/app/domain/model"
)

type ProgramReminderResp struct {
	ProgramName string `json:"program_name"`
	Code        string `json:"code"`
	model.ProgramCodeReminder
}

type EmaiReminderPayload struct {
	Name     string
	Code     string
	Deadline string
	Sample   int
}
