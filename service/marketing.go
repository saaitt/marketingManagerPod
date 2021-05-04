package service

import (
	"github.com/google/uuid"
	"github.com/saaitt/marketingManagerPod/model"
	"github.com/saaitt/marketingManagerPod/response"
)

type MarketingRepo interface {
	CreateProduct(item *model.MarketingProduct) error
	ListAllMarketingProducts(userId int) ([]model.MarketingProduct, error)
}

type MarketingService struct {
	Repo MarketingRepo
}

func (i MarketingService) CreateProduct(product model.Product, userId int) (*response.MarketingResponse, error) {

	marketingProduct := model.MarketingProduct{
		ProductId:  product.ID,
		Product:    product,
		UserID:     userId,
		UsageCount: 0,
		UUID:       uuid.New().String(),
	}
	if err := i.Repo.CreateProduct(&marketingProduct); err != nil {
		return nil, err
	}
	return &response.MarketingResponse{
		ID:         marketingProduct.ID,
		ProductId:  marketingProduct.ProductId,
		Product:    marketingProduct.Product,
		UserID:     marketingProduct.UserID,
		UsageCount: marketingProduct.UsageCount,
		UUID:       marketingProduct.UUID,
	}, nil
}

func (i MarketingService) ListAllProducts(userId int) ([]response.MarketingResponse, error) {
	items, err := i.Repo.ListAllMarketingProducts(userId)
	if err != nil {
		return nil, err
	}
	responses := []response.MarketingResponse{}
	for _, marketingProduct := range items {
		responses = append(responses, response.MarketingResponse{
			ID:         marketingProduct.ID,
			ProductId:  marketingProduct.ProductId,
			Product:    marketingProduct.Product,
			UserID:     marketingProduct.UserID,
			UsageCount: marketingProduct.UsageCount,
			UUID:       marketingProduct.UUID,
		})
	}
	return responses, nil
}
