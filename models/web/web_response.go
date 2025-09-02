package web

type WebResponse struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProductCategoryResponse struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
}

type ProductResponse struct {
	ID              int                     `json:"id"`
	Name            string                  `json:"name"`
	Slug            string                  `json:"slug"`
	Thumbnail       string                  `json:"thumbnail"`
	Price           int                     `json:"price"`
	Exp             string                  `json:"exp"`
	CategoryID      int                     `json:"category_id"`
	ProductCategory ProductCategoryResponse `json:"category"`
}