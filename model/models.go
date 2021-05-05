package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Product struct {
	gorm.Model
	ID       int
	Title    string
	PageLink string
	UserID   int
}

func (Product) TableName() string {
	return "products"
}

func (MarketingProduct) TableName() string {
	return "marketing_products"
}

type MarketingProduct struct {
	gorm.Model
	ID         int
	ProductId  int
	Product    Product `gorm:"foreignkey:id;references:product_id"`
	UserID     int
	UsageCount int
	UUID       string
}

func (User) TableName() string {
	return "users"
}

const (
	UserTypeAdmin    = "admin"
	UserTypeMarketer = "marketer"
)

type User struct {
	gorm.Model
	ID           int
	Username     string
	PasswordHash string
	UserType     string
}

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
