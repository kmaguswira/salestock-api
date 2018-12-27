package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Sku       string     `gorm:"type:varchar(255)" json:"sku,omitempty"`
	Name      string     `gorm:"type:varchar(255);unique" json:"name,omitempty"`
	Quantity   string    `gorm:"type:varchar(255)" json:"quantity,omitempty"`
}
