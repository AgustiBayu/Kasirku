package web

type ProductResponse struct {
	Id              int                     `json:"id"`
	Name            string                  `json:"name"`
	Slug            string                  `json:"slug"`
	Thumbnail       string                  `json:"thumbnail"`
	Price           int                     `json:"price"`
	Exp             string                  `json:"exp"`
	ProductCategory ProductCategoryResponse `json:"product_category"`
}

