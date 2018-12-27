package forms

type CreateProduct struct {
	Sku      string `json:"sku" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type UpdateProduct struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
