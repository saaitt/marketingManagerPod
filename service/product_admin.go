package service

import (
	"github.com/saaitt/marketingManagerPod/model"
	"github.com/saaitt/marketingManagerPod/request"
	"github.com/saaitt/marketingManagerPod/response"
)

type ProductRepo interface {
	Create(item *model.Product) error
	FindByUser(userID int) ([]model.Product, error)
	FindByID(id int) (*model.Product, error)
}

type ProductAdminService struct {
	Repo ProductRepo
}

func (i ProductAdminService) Create(request request.CreateProductRequest) (*response.ProductResponse, error) {
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

func (i ProductAdminService) FindByUser(userID int) ([]response.ProductResponse, error) {
	items, err := i.Repo.FindByUser(userID)
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
