package web

type ProductCategoryUpdateRequest struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
}
