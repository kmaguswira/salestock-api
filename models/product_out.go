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
	SellPrice  int     `gorm:"type:int" json:"sellPrice,omitempty"`
	TotalPrice int     `gorm:"type:int" json:"totalPrice,omitempty"`
	Type       string  `gorm:"type:int" json:"type,omitempty"`
	Note       string  `gorm:"type:varchar(255)" json:"note,omitempty"`
}

func (p *ProductOut) AfterCreate(db *gorm.DB) error {
	var product Product

	db.Where("ID = ?", p.ProductID).First(&product)

	product.Quantity -= p.Quantity
	db.Save(&product)

	return nil
}
