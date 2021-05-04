package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ID       int
	Title    string
	PageLink string
}

func (Product) TableName() string {
	return "products"
}

type MarketingProduct struct {
	gorm.Model
	ID         int
	ProductId  int     `gorm:"index"`
	Product    Product `gorm:"foreignkey:id;references:product_id"`
	UserID     int
	UsageCount int
	UUID       string
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

type SQLMarketingRepo struct {
	DB *gorm.DB
}

func (s SQLMarketingRepo) ListAllMarketingProducts(userId int) ([]MarketingProduct, error) {
	marketingProducts := []MarketingProduct{}
	if err := s.DB.Model(&MarketingProduct{}).Find(&marketingProducts, "UserID = ?", userId).Error; err != nil {
		return nil, err
	}
	return marketingProducts, nil
}

func (s SQLMarketingRepo) CreateProduct(marketingProduct *MarketingProduct) error {
	if err := s.DB.Create(&marketingProduct).Error; err != nil {
		return err
	}
	return nil
}
