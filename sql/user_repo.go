package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/saaitt/marketingManagerPod/model"
)

type UserRepo struct {
	DB *gorm.DB
}

func (s UserRepo) Create(user *model.User) error {
	if err := s.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s UserRepo) FindByUsername(username string) (*model.User, error) {
	user := model.User{}
	if err := s.DB.Model(&model.User{}).First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
