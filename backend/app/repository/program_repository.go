package repository

import (
	model "EQA/backend/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type ProgramRepository interface {
	FindAll() ([]model.Program, error)
}

type ProgramRepositoryImpl struct {
	db *gorm.DB
}

func seedDataProgram(db *gorm.DB) {
	if err := db.First(&model.Program{}, "name = ?", "TORCH").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Create(&model.Program{ID: 1, Name: "Vi sinh lâm sàng"})
			db.Create(&model.Program{ID: 2, Name: "TORCH"})
			db.Create(&model.Program{ID: 3, Name: "Syphilis"})
			db.Create(&model.Program{ID: 4, Name: "Huyết thanh học ký sinh trùng"})
			db.Create(&model.Program{ID: 5, Name: "Soi phân tìm kí sinh trùng đường ruột"})
			db.Create(&model.Program{ID: 6, Name: "HIV"})
			db.Create(&model.Program{ID: 7, Name: "Mycobacterium tuberculosis Real-time PCR (HCM207V)"})
			db.Create(&model.Program{ID: 8, Name: "Huyết thanh học viêm gan B, C"})
			db.Create(&model.Program{ID: 9, Name: "EBV"})
			db.Create(&model.Program{ID: 10, Name: "Vi sinh nhuộm soi"})
			db.Create(&model.Program{ID: 11, Name: "Vancomycin"})

			db.Create(&model.ProgramCode{ID: 1, Name: "CM16A", ProgramId: 1, Year: 2024})
			db.Create(&model.ProgramCode{ID: 2, Name: "TO06C", ProgramId: 2, Year: 2024})
			db.Create(&model.ProgramCode{ID: 3, Name: "SY06C", ProgramId: 3, Year: 2024})
			db.Create(&model.ProgramCode{ID: 4, Name: "PS05A", ProgramId: 4, Year: 2024})
			db.Create(&model.ProgramCode{ID: 5, Name: "QE1008 CK7", ProgramId: 5, Year: 2024})
			db.Create(&model.ProgramCode{ID: 6, Name: "P.I 24", ProgramId: 6, Year: 2024})
			db.Create(&model.ProgramCode{ID: 7, Name: "QE1016 QTT0178", ProgramId: 7, Year: 2024})
			db.Create(&model.ProgramCode{ID: 8, Name: "HS05A", ProgramId: 8, Year: 2024})
			db.Create(&model.ProgramCode{ID: 9, Name: "EB02C", ProgramId: 9, Year: 2024})
			db.Create(&model.ProgramCode{ID: 10, Name: "EB03C", ProgramId: 9, Year: 2024})
			db.Create(&model.ProgramCode{ID: 11, Name: "QE1021 CK5", ProgramId: 10, Year: 2024})
			db.Create(&model.ProgramCode{ID: 12, Name: "IM15C", ProgramId: 11, Year: 2024})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 1, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 1, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 1, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 5, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 6, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 7, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 8, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 9, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 10, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 11, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 12, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 13, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 14, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 15, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 16, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 17, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 18, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 19, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 20, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 21, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 22, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 23, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 2, Sample: 24, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 5, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 6, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 7, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 8, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 9, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 10, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 11, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 3, Sample: 12, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 4, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 4, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 4, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 5, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 5, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 5, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 5, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 6, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 6, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 7, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 7, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 7, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 8, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 8, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 8, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 8, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 9, Sample: 7, IsDefault: true, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 9, Sample: 8, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 9, Sample: 9, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 9, Sample: 10, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 9, Sample: 11, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 9, Sample: 12, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 10, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 10, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 10, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 10, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 10, Sample: 5, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 10, Sample: 6, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 11, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 11, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 11, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 11, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 1, IsDefault: true, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 5, IsDefault: true, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 6, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 7, IsDefault: true, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 8, IsDefault: true, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 9, IsDefault: true, DateOfReturn: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 10, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 11, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeReminder{ProgramCodeId: 12, Sample: 12, IsDefault: true, DateOfReturn: time.Date(2024, time.December, 1, 12, 0, 0, 0, time.UTC)})
		}
	}
}

func ProgramRepositoryInit(db *gorm.DB) *ProgramRepositoryImpl {
	db.AutoMigrate(&model.Program{})
	db.AutoMigrate(&model.ProgramCode{})
	db.AutoMigrate(&model.ProgramCodeReminder{})
	seedDataProgram(db)
	return &ProgramRepositoryImpl{
		db: db,
	}
}

func (u ProgramRepositoryImpl) FindAll() ([]model.Program, error) {
	var data []model.Program
	err := u.db.Model(&model.Program{}).Preload("ProgramCodes.ProgramCodeReminders").Order("id ASC").Find(&data).Error
	if err != nil {
		return []model.Program{}, err
	}
	return data, nil
}
