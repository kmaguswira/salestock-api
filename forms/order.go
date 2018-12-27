package forms

import (
	"time"
)

type CreateOrder struct {
	ProductID     int       `json:"productId" binding:"required"`
	OrderQuantity int       `json:"orderQuantity" binding:"required"`
	BasePrice     int       `json:"basePrice" binding:"required"`
	Invoice       string    `json:"invoice"`
	Date          time.Time `json:"date"`
}

type UpdateOrder struct {
	ProductID     int       `json:"productId"`
	OrderQuantity int       `json:"orderQuantity"`
	BasePrice     int       `json:"basePrice"`
	TotalPrice    int       `json:"totalPrice"`
	Invoice       string    `json:"invoice"`
	Date          time.Time `json:"date"`
}
