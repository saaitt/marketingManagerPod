package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/saaitt/marketingManagerPod/model"
)

type ProductRepo struct {
	DB *gorm.DB
}

func (s ProductRepo) Create(product *model.Product) error {
	if err := s.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (s ProductRepo) ListAll() ([]model.Product, error) {
	products := []model.Product{}
	if err := s.DB.Model(&model.Product{}).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (s ProductRepo) FindOne(id int) ([]model.Product, error) {
	product := []model.Product{}
	if err := s.DB.Model(&model.Product{}).First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return product, nil
}