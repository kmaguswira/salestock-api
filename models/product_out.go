package models

import (
	"github.com/jinzhu/gorm"
)

type ProductOut struct {
	gorm.Model
	Product    Product `json:"product,omitempty"`
	ProductID  uint    `json:"productId,omitempty"`
	Sales      Sales   `json:"sales,omitempty"`
	SalesID    uint    `json:"salesId,omitempty"`
	Quantity   int     `gorm:"type:int" json:"quantity,omitempty"`
	SellPrice  int     `gorm:"type:int;default:'0'" json:"sellPrice"`
	TotalPrice int     `gorm:"type:int;default:'0'" json:"totalPrice"`
	Type       string  `gorm:"type:int" json:"type,omitempty"`
	Note       string  `gorm:"type:varchar(255)" json:"note,omitempty"`
}
