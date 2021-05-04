package service

import (
	"github.com/google/uuid"
	"github.com/saaitt/marketingManagerPod/model"
	"github.com/saaitt/marketingManagerPod/response"
)

type MarketingRepo interface {
	CreateProduct(item *model.MarketingProduct) error
	ListAllMarketingProductsByUserID(userId int) ([]model.MarketingProduct, error)
	FindProduct(uuid string) (string, error)
	IncreaseUrlUsage(uuid string) error
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

func (i MarketingService) ListAllMarketingProductsByUserID(userId int) ([]response.MarketingResponse, error) {
	marketingProducts, err := i.Repo.ListAllMarketingProductsByUserID(userId)
	if err != nil {
		return nil, err
	}
	responses := []response.MarketingResponse{}
	for _, marketingProduct := range marketingProducts {
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
func (i MarketingService) ResolvePage(uuid string) (string, error) {
	PageLink, err := i.Repo.FindProduct(uuid)
	if err != nil {
		return "", err
	}
	err = i.Repo.IncreaseUrlUsage(uuid)
	if err != nil {
		return "", err
	}
	return PageLink, nil
}
