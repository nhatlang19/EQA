package dto

import (
	model "EQA/backend/app/domain/model"
)

type ProgramReminderResp struct {
	ProgramName string `json:"program_name"`
	Code        string `json:"code"`
	model.ProgramCodeDetail
}

type ProgramExportResp struct {
	ProgramName  string `json:"program_name"`
	Code         string `json:"code"`
	ProviderName string `json:"provider_name"`
	model.ProgramCodeDetail
}

type EmaiReminderPayload struct {
	Name     string
	Code     string
	Deadline string
	Sample   int
}

type UpdateProgramResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateProgramCodeResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

type CreateProgramCodeResp struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

type ProgramCodeDetailReq struct {
	ProgramCodeDetails []model.ProgramCodeDetail `json:"program_code_details"`
}

type ProgramCodeFilter struct {
	Year int `json:"year"`
}
