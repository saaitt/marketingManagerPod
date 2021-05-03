package service

import(
	"github.com/saaitt/marketingManagerPod/model"
	"github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/response"

)

type Product interface {
	Create(item *model.Product) error
	ListAll() ([]model.Product, error)
}

type ProductService struct {
	Repo Product
}

func (i ProductService) Create(request request.CreateProductRequest) (*response.ProductResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	product := model.Product{
		Title: request.Title,
		PageLink:  request.PageLink,
	}
	if err := i.Repo.Create(&product); err != nil {
		return nil, err
	}
	return &response.ProductResponse{
		ID:    product.ID,
		Title: product.Title,
		PageLink:  request.PageLink,
	}, nil
}

func (i ProductService) ListAll() ([]response.ProductResponse, error) {
	items, err := i.Repo.ListAll()
	if err != nil {
		return nil, err
	}
	responses := []response.ProductResponse{}
	for _, product := range items {
		responses = append(responses, response.ProductResponse{
			ID:    product.ID,
			Title: product.Title,
			PageLink:  product.PageLink,
		})
	}
	return responses, nil
}
