package response

import (
	"github.com/saaitt/marketingManagerPod/model"
)

type ProductResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	PageLink string `json:"pageLink"`
}

type MarketingResponse struct {
	ID         int           `json:"id"`
	ProductId  int           `json:"product_id"`
	Product    model.Product `json:"product"`
	UserID     int           `json:"user_id"`
	UsageCount int           `json:"usage_count"`
	UUID       string        `json:"uuid"`
}
