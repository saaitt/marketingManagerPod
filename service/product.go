package service

import (
	"github.com/saaitt/marketingManagerPod/model"
	"github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/response"
)

type ProductRepo interface {
	Create(item *model.Product) error
	ListAll() ([]model.Product, error)
	FindOne(id int) ([]model.Product, error)
}

type ProductService struct {
	Repo ProductRepo
}

func (i ProductService) Create(request request.CreateProductRequest) (*response.ProductResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	product := model.Product{
		UserID:   request.UserID,
		Title:    request.Title,
		PageLink: request.PageLink,
	}
	if err := i.Repo.Create(&product); err != nil {
		return nil, err
	}
	return &response.ProductResponse{
		ID:       product.ID,
		UserID:   product.UserID,
		Title:    product.Title,
		PageLink: request.PageLink,
	}, nil
}

func (i ProductService) FindByUser(int) ([]response.ProductResponse, error) {
	items, err := i.Repo.ListAll()
	if err != nil {
		return nil, err
	}
	responses := []response.ProductResponse{}
	for _, product := range items {
		responses = append(responses, response.ProductResponse{
			ID:       product.ID,
			Title:    product.Title,
			PageLink: product.PageLink,
		})
	}
	return responses, nil
}
func (i ProductService) FindOne(id int) ([]response.ProductResponse, error) {
	items, err := i.Repo.FindOne(id)
	if err != nil {
		return nil, err
	}
	responses := []response.ProductResponse{}
	for _, product := range items {
		responses = append(responses, response.ProductResponse{
			ID:       product.ID,
			Title:    product.Title,
			PageLink: product.PageLink,
		})
	}
	return responses, nil
}
