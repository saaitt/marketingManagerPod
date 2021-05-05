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

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

func (c CreateUserRequest) Validate() error {
	if c.Username == "" {
		return errors.New("username id is required")
	}
	if c.Password == "" {
		return errors.New("password id is required")
	}
	if c.UserType == "" {
		return errors.New("user type id is required")
	}
	return nil
}

type LoggedInUserRequest struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

func (c LoggedInUserRequest) Validate() error {
	if c.Username == "" {
		return errors.New("username id is required")
	}
	if c.PasswordHash == "" {
		return errors.New("password hash id is required")
	}
	return nil
}
