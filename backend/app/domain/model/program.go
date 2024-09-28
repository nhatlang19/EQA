package model

import "time"

type Program struct {
	ID           int           `gorm:"column:id; primary_key; not null" json:"id"`
	Name         string        `gorm:"column:name;" json:"name"`
	ProgramCodes []ProgramCode `gorm:"foreignkey:ProgramId" json:"program_codes"`
	ProviderId   int           `gorm:"column:provider_id;" json:"provider_id"`
	Provider     Provider      `gorm:"foreignkey:ProviderId" json:"provider"`
	BaseModel
}

type ProgramCode struct {
	ID                   int                   `gorm:"column:id; primary_key;not null;" sql:"AUTO_INCREMENT" json:"id"`
	Name                 string                `gorm:"column:name;" json:"name"`
	ProgramId            int                   `gorm:"column:program_id;" json:"program_id"`
	Year                 int                   `gorm:"column:year;" json:"year"`
	ProgramCodeReminders []ProgramCodeReminder `gorm:"foreignkey:ProgramCodeId;omitempty" json:"program_code_reminders"`
	Status               int                   `gorm:"column:status;default:1" json:"status"`
	BaseModel
}

type ProgramCodeReminder struct {
	ID            int       `gorm:"column:id; primary_key;not null" json:"id"`
	ProgramCodeId int       `gorm:"column:program_code_id;" json:"program_code_id"`
	Sample        int       `gorm:"column:sample;" json:"sample"`
	DateOfReceive time.Time `gorm:"column:date_of_receive;default:CURRENT_TIMESTAMP" json:"date_of_receive"`
	DateOfReturn  time.Time `gorm:"column:date_of_return;" json:"date_of_return"`
	IsDefault     bool      `gorm:"column:is_default;default: false" json:"is_default"`
	Status        int       `gorm:"column:status;default: 0" json:"status"`
	PercentPassed int       `gorm:"column:percent_passed;" json:"percent_passed"`
	BaseModel
}
