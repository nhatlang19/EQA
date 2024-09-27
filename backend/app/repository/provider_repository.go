package repository

import (
	model "EQA/backend/app/domain/model"

	"gorm.io/gorm"
)

type ProviderRepository interface {
}

type ProviderRepositoryImpl struct {
	db *gorm.DB
}

func seedDataProvider(db *gorm.DB) {
	if err := db.First(&model.Provider{}, "name = ?", "Pasteur").Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			db.Create(&model.Provider{ID: 1, Name: "Trung tâm kiểm chuẩn xét nghiệm thành phố Hồ Chí Minh"})
			db.Create(&model.Provider{ID: 2, Name: "Trung tâm kiểm chuẩn Chất lượng xét nghiệm Y học - Đại học Y Dược TP Hồ Chí Minh"})
			db.Create(&model.Provider{ID: 3, Name: "Pasteur"})
		}
	}
}

func ProviderRepositoryInit(db *gorm.DB) *ProviderRepositoryImpl {
	db.AutoMigrate(&model.Provider{})
	seedDataProvider(db)
	return &ProviderRepositoryImpl{
		db: db,
	}
}
