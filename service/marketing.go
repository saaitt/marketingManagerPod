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
	MarketingRepo MarketingRepo
	ProductRepo   ProductRepo
}

func (i MarketingService) CreateProduct(productId int, userId int) (*response.MarketingResponse, error) {
	product, err := i.ProductRepo.FindOne(productId)
	if err != nil {
		return nil, err
	}

	marketingProduct := model.MarketingProduct{
		ProductId:  product[0].ID,
		Product:    product[0],
		UserID:     userId,
		UsageCount: 0,
		UUID:       uuid.New().String(),
	}

	if err := i.MarketingRepo.CreateProduct(&marketingProduct); err != nil {
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
	marketingProducts, err := i.MarketingRepo.ListAllMarketingProductsByUserID(userId)
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
	PageLink, err := i.MarketingRepo.FindProduct(uuid)
	if err != nil {
		return "", err
	}
	err = i.MarketingRepo.IncreaseUrlUsage(uuid)
	if err != nil {
		return "", err
	}
	return PageLink, nil
}
