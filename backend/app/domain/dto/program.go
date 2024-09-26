package dto

import (
	model "EQA/backend/app/domain/model"
)

type ProgramResp struct {
	ID           int                 `gorm:"column:id; primary_key; not null" json:"id"`
	Name         string              `gorm:"column:name;" json:"name"`
	ProgramCodes []model.ProgramCode `gorm:"foreignkey:ProgramId" json:"program_codes"`
}
