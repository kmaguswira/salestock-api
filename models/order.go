package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Product       Product         `json:"product,omitempty"`
	ProductID     uint            `json:"productId,omitempty"`
	OrderQuantity int             `gorm:"type:int" json:"orderQuantity,omitempty"`
	BasePrice     int             `gorm:"type:int" json:"basePrice,omitempty"`
	TotalPrice    int             `gorm:"type:int" json:"totalPrice,omitempty"`
	Status        string          `gorm:"type:varchar(255);default:'Incomplete'" json:"status,omitempty"`
	Invoice       string          `gorm:"type:varchar(255)" json:"invoice,omitempty"`
	Date          time.Time       `json:"date,omitempty"`
	OrderProgress []OrderProgress `json:"orderProgress"`
}
