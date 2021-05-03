package model

import "github.com/jinzhu/gorm"

type Product struct {
	ID    int
	Title string
	PageLink  string
}

func (Product) TableName() string {
	return "products"
}

type MarketingProduct struct {
	ID    int
	ProductID  int
	UserID int
	CountUsage int
	UUID string
}

func (MarketingProduct) TableName() string {
	return "marketing_products"
}



type SQLItemRepo struct {
	DB *gorm.DB
}


func (s SQLItemRepo) Create(product *Product) error {
	if err := s.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (s SQLItemRepo) ListAll() ([]Product, error) {
	products := []Product{}
	if err := s.DB.Model(&Product{}).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}


