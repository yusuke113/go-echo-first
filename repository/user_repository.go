package repository

import (
	"asobi/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

// ユーザーを登録
func (ur *userRepository) Create(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
