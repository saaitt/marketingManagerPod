package service

import(
	"github.com/saaitt/marketingManagerPod/model"
	// "github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/response"
	"github.com/google/uuid"
)

type MarketingRepo interface {
	CreateProduct(item *model.MarketingProduct) error
	ListAllMarketingProducts(userId int) ([]model.MarketingProduct, error)
}

type MarketingService struct{
	Repo MarketingRepo
}


func (i MarketingService) CreateProduct(product model.Product,userId int) (*response.MarketingResponse, error) {
	
	marketingProduct := model.MarketingProduct{
	ProductId : product.ID,
	Product   : product,
	UserID    : userId,
	CountUsage: 0,
	UUID      :uuid.New().String(),
	}
	if err := i.Repo.CreateProduct(&marketingProduct); err != nil {
		return nil, err
	}
	return &response.MarketingResponse{
		MarketingProduct: marketingProduct,
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
			MarketingProduct: marketingProduct,
		})
	}
	return responses, nil
}
