package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/saaitt/marketingManagerPod/model"
)

type MarketingRepo struct {
	DB *gorm.DB
}

func (s MarketingRepo) ListAllMarketingProductsByUserID(userId int) ([]model.MarketingProduct, error) {
	marketingProducts := []model.MarketingProduct{}
	if err := s.DB.Model(&model.MarketingProduct{}).Find(&marketingProducts, "UserID = ?", userId).Error; err != nil {
		return nil, err
	}
	return marketingProducts, nil
}

func (s MarketingRepo) CreateProduct(marketingProduct *model.MarketingProduct) error {
	if err := s.DB.Create(&marketingProduct).Error; err != nil {
		return err
	}
	return nil
}

func (s MarketingRepo) FindProduct(uuid string) (string, error) {
	product := model.Product{}
	if err := s.DB.Model(&model.Product{}).Select("products.id, products.title, products.page_url").Joins("inner join marketing_products on marketing_products.productId = products.id").Where("marketing_products.uuid = ?", uuid).Find(&product).Error; err != nil {
		return "", err
	}
	return product.PageLink, nil
}

func (s MarketingRepo) IncreaseUrlUsage(uuid string) error {
	if err := s.DB.Exec("UPDATE marketing_products SET UsageCount = UsageCount + 1 WHERE uuid = ?", uuid).Error; err != nil {
		return err
	}
	return nil
}
