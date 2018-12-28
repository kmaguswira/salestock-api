package models

import (
	"github.com/jinzhu/gorm"
)

type OrderProgress struct {
	gorm.Model
	Order            Order `json:"order,omitempty"`
	OrderID          uint  `json:"orderId,omitempty"`
	QuantityReceived int   `gorm:"type:int" json:"quantityReceived,omitempty"`
}

func (o *OrderProgress) AfterCreate(db *gorm.DB) error {
	var order Order
	var product Product

	db.Where("ID = ?", o.OrderID).First(&order)
	db.Where("ID = ?", order.ProductID).First(&product)

	product.Quantity += o.QuantityReceived
	db.Save(&product)

	return nil
}

func (o *OrderProgress) BeforeUpdate(db *gorm.DB) error {
	var order Order
	var orderProgress OrderProgress
	var product Product

	db.Where("ID = ?", o.ID).First(&orderProgress)
	db.Where("ID = ?", o.OrderID).First(&order)
	db.Where("ID = ?", order.ProductID).First(&product)

	product.Quantity -= orderProgress.QuantityReceived
	db.Save(&product)

	return nil
}

func (o *OrderProgress) AfterUpdate(db *gorm.DB) error {
	var order Order
	var product Product

	db.Where("ID = ?", o.OrderID).First(&order)
	db.Where("ID = ?", order.ProductID).First(&product)

	product.Quantity += o.QuantityReceived
	db.Save(&product)

	return nil
}

func (o *OrderProgress) AfterDelete(db *gorm.DB) error {
	var order Order
	var product Product

	db.Where("ID = ?", o.OrderID).First(&order)
	db.Where("ID = ?", order.ProductID).First(&product)

	product.Quantity -= o.QuantityReceived
	db.Save(&product)

	return nil
}

func (o *OrderProgress) AfterSave(db *gorm.DB) error {
	var orderProgress []OrderProgress
	var order Order
	total := 0

	db.Where("ID = ?", o.OrderID).First(&order)
	db.Where("order_id = ?", o.OrderID).Find(&orderProgress)

	for _, op := range orderProgress {
		total += op.QuantityReceived
	}

	if total >= order.OrderQuantity {
		order.Status = "Complete"
	} else {
		order.Status = "Incomplete"
	}

	db.Save(&order)

	return nil
}
