package response

type ProductResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	PageLink string `json:"pageLink"`
}
