package request

import "errors"

type CreateProductRequest struct {
	Title    string `json:"title"`
	ID       int    `json:"id"`
	PageLink string `json:"pageLink"`
}

func (c CreateProductRequest) Validate() error {
	if c.Title == "" {
		return errors.New("title is required")
	}
	if c.PageLink == "" {
		return errors.New("page link is required")
	}
	return nil
}

type CreateMarketingRequest struct {
	ProductId int `json:"product_id"`
	UserID    int `json:"user_id"`
}

func (c CreateMarketingRequest) Validate() error {
	if c.ProductId == 0 {
		return errors.New("product id is required")
	}
	if c.UserID == 0 {
		return errors.New("user id is required")
	}
	return nil
}
