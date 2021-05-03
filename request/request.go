package request

import "errors"

type CreateProductRequest struct {
	Title    string `json:"title"`
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
