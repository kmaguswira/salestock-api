package forms

type CreateProductOut struct {
	ProductID  uint   `json:"productId" binding:"required"`
	SalesID    uint   `json:"salesId"`
	Quantity   int    `json:"quantity" binding:"required"`
	SellPrice  int    `json:"sellPrice"`
	TotalPrice int    `json:"totalPrice"`
	Type       string `json:"type" binding:"required"`
	Note       string `json:"note"`
}

type UpdateProductOut struct {
	ProductID  uint   `json:"productId"`
	SalesID    uint   `json:"salesId"`
	Quantity   int    `json:"quantity"`
	SellPrice  int    `json:"sellPrice"`
	TotalPrice int    `json:"totalPrice"`
	Type       string `json:"type"`
	Note       string `json:"note"`
}
