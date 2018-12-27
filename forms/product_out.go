package forms

type CreateProductOut struct {
	ProductID  int    `json:"productId" binding:"required"`
	SalesID    int    `json:"salesId" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required"`
	SellPrice  int    `json:"sellPrice" binding:"required"`
	TotalPrice int    `json:"totalPrice" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Note       string `json:"note"`
}

type UpdateProductOut struct {
	ProductID  int    `json:"productId"`
	SalesID    int    `json:"salesId"`
	Quantity   int    `json:"quantity"`
	SellPrice  int    `json:"sellPrice"`
	TotalPrice int    `json:"totalPrice"`
	Type       string `json:"type"`
	Note       string `json:"note"`
}
