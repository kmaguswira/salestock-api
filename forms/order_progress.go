package forms

type CreateOrderProgress struct {
	OrderID          uint `json:"orderId" binding:"required"`
	QuantityReceived int  `json:"quantityReceived" binding:"required"`
}

type UpdateOrderProgress struct {
	QuantityReceived int `json:"quantityReceived"`
}
