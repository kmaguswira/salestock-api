package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type OrderProgress struct {
	gorm.Model
	Order            Order     `json:"order,omitempty"`
	OrderID          uint      `json:"orderId,omitempty"`
	QuantityReceived int       `gorm:"type:int" json:"quantityReceived,omitempty"`
	Date             time.Time `json:"date,omitempty"`
}
