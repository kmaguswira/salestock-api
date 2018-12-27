package forms

import (
	"time"
)

type CreateOrderProgress struct {
	OrderID          int       `json:"orderId" binding:"required"`
	QuantityReceived int       `json:"quantityReceived" binding:"required"`
	Date             time.Time `json:"date" binding:"required"`
}

type UpdateOrderProgress struct {
	OrderID          int       `json:"orderId"`
	QuantityReceived int       `json:"quantityReceived"`
	Date             time.Time `json:"date"`
}
