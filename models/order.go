package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Product       Product   `json:"product,omitempty"`
	ProductID     int       `json:"productId,omitempty"`
	OrderQuantity int       `gorm:"type:int" json:"orderQuantity,omitempty"`
	BasePrice     int       `gorm:"type:int" json:"basePrice,omitempty"`
	TotalPrice    int       `gorm:"type:int" json:"totalPrice,omitempty"`
	Status        int       `gorm:"type:varchar(255)" json:"status,omitempty"`
	Invoice       string    `gorm:"type:varchar(255)" json:"invoice,omitempty"`
	Date          time.Time `json:"date,omitempty"`
}
