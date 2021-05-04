package response

import(

	"github.com/saaitt/marketingManagerPod/model"

)
type ProductResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	PageLink string `json:"pageLink"`
}

type MarketingResponse struct {
	MarketingProduct model.MarketingProduct	`json:"array"`
}