package repository

import (
	model "EQA/backend/app/domain/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&model.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := u.db.First(&user, "username = ? AND active = ?", username, true).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
