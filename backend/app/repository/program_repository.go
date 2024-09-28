package repository

import (
	dto "EQA/backend/app/domain/dto"
	model "EQA/backend/app/domain/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ProgramRepository interface {
	FindAll() ([]model.Program, error)
	FindForExport() ([]dto.ProgramExportResp, error)
	FindOne(id int) (model.Program, error)
	FindOneWithFilter(id int, filter dto.ProgramCodeFilter) (model.Program, error)
	FindForReminder(num int) ([]dto.ProgramReminderResp, error)
	Save(program *model.Program) (model.Program, error)

	FindByCodeId(id int, codeId int) (model.ProgramCode, error)
	SaveCode(programCode *model.ProgramCode) (model.ProgramCode, error)
	CreateCode(programCode *model.ProgramCode) (model.ProgramCode, error)
	DeleteCodeById(id int) error

	SaveCodeDetail(data *model.ProgramCodeDetail) (model.ProgramCodeDetail, error)
	DeleteDetailCodeById(id int) error
}

type ProgramRepositoryImpl struct {
	db *gorm.DB
}

func seedDataProgram(db *gorm.DB) {
	if err := db.First(&model.Program{}, "name = ?", "TORCH").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Create(&model.Program{ID: 1, Name: "Vi sinh lâm sàng", ProviderId: 1})
			db.Create(&model.Program{ID: 2, Name: "TORCH", ProviderId: 1})
			db.Create(&model.Program{ID: 3, Name: "Syphilis", ProviderId: 1})
			db.Create(&model.Program{ID: 4, Name: "Huyết thanh học ký sinh trùng", ProviderId: 1})
			db.Create(&model.Program{ID: 5, Name: "Soi phân tìm kí sinh trùng đường ruột", ProviderId: 2})
			db.Create(&model.Program{ID: 6, Name: "Huyết thanh học HIV", ProviderId: 3})
			db.Create(&model.Program{ID: 7, Name: "Mycobacterium tuberculosis Real-time PCR", ProviderId: 2})
			db.Create(&model.Program{ID: 8, Name: "Huyết thanh học viêm gan B, C", ProviderId: 1})
			db.Create(&model.Program{ID: 9, Name: "EBV", ProviderId: 1})
			db.Create(&model.Program{ID: 10, Name: "Vi sinh nhuộm soi", ProviderId: 2})
			db.Create(&model.Program{ID: 11, Name: "Vancomycin", ProviderId: 2})
			db.Exec("SELECT setval('programs_id_seq', (SELECT MAX(id) FROM programs));")

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
			db.Create(&model.ProgramCode{ID: 12, Name: "IM16C", ProgramId: 11, Year: 2024})
			db.Exec("SELECT setval('program_codes_id_seq', (SELECT MAX(id) FROM program_codes));")

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 1, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 1, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 1, Sample: 3, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 1, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 3, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 4, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 5, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 6, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 7, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 8, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 9, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 10, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 11, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 12, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 13, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 14, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 15, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 16, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 17, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.September, 30, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 18, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.September, 30, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 19, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 20, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 21, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 22, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 23, IsDefault: true, DateOfReturn: time.Date(2024, time.December, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 2, Sample: 24, IsDefault: true, DateOfReturn: time.Date(2024, time.December, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 3, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 4, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 5, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 6, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 7, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 8, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 9, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 10, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 11, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 3, Sample: 12, IsDefault: true, DateOfReturn: time.Date(2024, time.December, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 4, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 4, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 4, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.September, 30, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 5, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 5, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 5, Sample: 3, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 5, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 6, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 6, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 7, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 7, Sample: 2, IsDefault: true, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 7, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 8, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 8, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 8, Sample: 3, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 8, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 9, Sample: 7, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 9, Sample: 8, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 9, Sample: 9, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 9, Sample: 10, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 9, Sample: 11, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 9, Sample: 12, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 10, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 10, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 10, Sample: 3, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 10, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 10, Sample: 5, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 10, Sample: 6, IsDefault: true, DateOfReturn: time.Date(2024, time.December, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 11, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 11, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 11, Sample: 3, IsDefault: true, DateOfReturn: time.Date(2024, time.September, 30, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 11, Sample: 4, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})

			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 1, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.January, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 2, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.February, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 3, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.March, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 4, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.April, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 5, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.May, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 6, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 7, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.July, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 8, IsDefault: true, Status: 1, PercentPassed: 100, DateOfReturn: time.Date(2024, time.August, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 9, IsDefault: true, DateOfReturn: time.Date(2024, time.September, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 10, IsDefault: true, DateOfReturn: time.Date(2024, time.October, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 11, IsDefault: true, DateOfReturn: time.Date(2024, time.November, 1, 12, 0, 0, 0, time.UTC)})
			db.Create(&model.ProgramCodeDetail{ProgramCodeId: 12, Sample: 12, IsDefault: true, DateOfReturn: time.Date(2024, time.December, 1, 12, 0, 0, 0, time.UTC)})
		}
	}
}

func ProgramRepositoryInit(db *gorm.DB) *ProgramRepositoryImpl {
	db.AutoMigrate(&model.Provider{})
	db.AutoMigrate(&model.Program{})
	db.AutoMigrate(&model.ProgramCode{})
	db.AutoMigrate(&model.ProgramCodeDetail{})
	seedDataProgram(db)
	return &ProgramRepositoryImpl{
		db: db,
	}
}

func (u ProgramRepositoryImpl) FindAll() ([]model.Program, error) {
	var data []model.Program
	currentTime := time.Now()
	err := u.db.Model(&model.Program{}).Preload("Provider").
		Preload("ProgramCodes", func(db *gorm.DB) *gorm.DB {
			return db.Where("year = ?", currentTime.Year())
		}).
		Preload("ProgramCodes.ProgramCodeDetails").
		Order("id ASC").Find(&data).Error
	if err != nil {
		return []model.Program{}, err
	}
	return data, nil
}

func (u ProgramRepositoryImpl) FindOne(id int) (model.Program, error) {
	result := model.Program{
		ID: id,
	}

	err := u.db.
		Preload("Provider", func(db *gorm.DB) *gorm.DB {
			return db.Order("id desc")
		}).
		Preload("ProgramCodes", func(db *gorm.DB) *gorm.DB {
			return db.Order("id desc")
		}).
		Preload("ProgramCodes.ProgramCodeDetails", func(db *gorm.DB) *gorm.DB {
			return db.Order("id desc")
		}).
		Debug().
		First(&result).Error
	if err != nil {
		return model.Program{}, err
	}
	return result, nil
}

func (u ProgramRepositoryImpl) FindOneWithFilter(id int, filter dto.ProgramCodeFilter) (model.Program, error) {
	result := model.Program{
		ID: id,
	}

	err := u.db.
		Preload("Provider", func(db *gorm.DB) *gorm.DB {
			return db.Order("id desc")
		}).
		Preload("ProgramCodes", func(db *gorm.DB) *gorm.DB {
			if filter.Year != -1 {
				return db.Where("year = ?", filter.Year).Order("id desc")
			}
			return db.Order("id desc")
		}).
		Preload("ProgramCodes.ProgramCodeDetails", func(db *gorm.DB) *gorm.DB {
			return db.Order("sample asc")
		}).
		Debug().
		First(&result).Error
	if err != nil {
		return model.Program{}, err
	}
	return result, nil
}

func (u ProgramRepositoryImpl) FindForReminder(num int) ([]dto.ProgramReminderResp, error) {
	var data []dto.ProgramReminderResp

	err := u.db.Model(&model.Program{}).
		Joins("LEFT JOIN program_codes as program_codes ON program_codes.program_id = programs.id").
		Joins("LEFT JOIN program_code_details as program_code_details ON program_codes.id = program_code_details.program_code_id").
		Where("program_code_details.date_of_return::DATE - NOW()::DATE <= ?", num).
		Where("program_code_details.date_of_return::DATE - NOW()::DATE > ?", 0).
		Select("programs.name as program_name, program_codes.name as code, program_code_details.*").
		Find(&data).Error
	if err != nil {
		return []dto.ProgramReminderResp{}, err
	}
	return data, nil
}

func (u ProgramRepositoryImpl) FindForExport() ([]dto.ProgramExportResp, error) {
	var data []dto.ProgramExportResp
	currentTime := time.Now()
	err := u.db.Model(&model.Program{}).
		Joins("LEFT JOIN providers as providers ON programs.provider_id = providers.id").
		Joins("LEFT JOIN program_codes as program_codes ON program_codes.program_id = programs.id").
		Joins("LEFT JOIN program_code_details as program_code_details ON program_codes.id = program_code_details.program_code_id").
		Where("program_codes.year = ?", currentTime.Year()).
		Where("program_code_details.status = ?", 1).
		Select("programs.name as program_name, providers.name as provider_name, program_codes.name as code, program_code_details.*").
		Debug().
		Find(&data).Error
	if err != nil {
		return []dto.ProgramExportResp{}, err
	}
	return data, nil
}

func (u ProgramRepositoryImpl) Save(data *model.Program) (model.Program, error) {
	var err error
	err = u.db.Save(data).Error
	if err != nil {
		fmt.Println(err)
		return model.Program{}, err
	}

	return *data, nil
}

func (u ProgramRepositoryImpl) FindByCodeId(id int, codeId int) (model.ProgramCode, error) {
	result := model.ProgramCode{
		ID:        codeId,
		ProgramId: id,
	}
	err := u.db.Model(&model.ProgramCode{}).First(&result).Error
	if err != nil {
		return model.ProgramCode{}, err
	}
	return result, nil
}

func (u ProgramRepositoryImpl) SaveCode(data *model.ProgramCode) (model.ProgramCode, error) {
	var err error
	err = u.db.Save(data).Error
	if err != nil {
		fmt.Println(err)
		return model.ProgramCode{}, err
	}

	return *data, nil
}

func (u ProgramRepositoryImpl) CreateCode(data *model.ProgramCode) (model.ProgramCode, error) {
	var err error
	err = u.db.Omit("ID").Create(data).Error
	if err != nil {
		fmt.Println(err)
		return model.ProgramCode{}, err
	}

	return *data, nil
}

func (u ProgramRepositoryImpl) SaveCodeDetail(data *model.ProgramCodeDetail) (model.ProgramCodeDetail, error) {
	var err error
	err = u.db.Save(data).Error
	if err != nil {
		fmt.Println(err)
		return model.ProgramCodeDetail{}, err
	}

	return *data, nil
}

func (u ProgramRepositoryImpl) DeleteCodeById(id int) error {
	err := u.db.Where("id = ?", id).Delete(&model.ProgramCode{}).Error
	if err != nil {
		fmt.Println("Got an error when delete program code. Error: ", err)
		return err
	}
	return nil
}

func (u ProgramRepositoryImpl) DeleteDetailCodeById(id int) error {
	err := u.db.Where("program_code_id = ?", id).Delete(&model.ProgramCodeDetail{}).Error
	if err != nil {
		fmt.Println("Got an error when delete program code detail. Error: ", err)
		return err
	}
	return nil
}
