package forms

type CreateOrder struct {
	ProductID     uint      `json:"productId" binding:"required"`
	OrderQuantity int       `json:"orderQuantity" binding:"required"`
	BasePrice     int       `json:"basePrice" binding:"required"`
	Invoice       string    `json:"invoice"`
}

type UpdateOrder struct {
	ProductID     uint      `json:"productId"`
	OrderQuantity int       `json:"orderQuantity"`
	BasePrice     int       `json:"basePrice"`
	TotalPrice    int       `json:"totalPrice"`
	Invoice       string    `json:"invoice"`
}
