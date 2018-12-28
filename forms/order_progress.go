package forms

import (
	"time"
)

type CreateOrderProgress struct {
	OrderID          uint      `json:"orderId" binding:"required"`
	QuantityReceived int       `json:"quantityReceived" binding:"required"`
	Date             time.Time `json:"date" binding:"required"`
}

type UpdateOrderProgress struct {
	OrderID          uint      `json:"orderId"`
	QuantityReceived int       `json:"quantityReceived"`
	Date             time.Time `json:"date"`
}
