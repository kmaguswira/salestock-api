package utils

import (
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/models"
	"reflect"
	"strconv"
	"strings"
)

func CopyValue(a, b interface{}) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b).Elem()

	at := av.Type()

	for i := 0; i < at.NumField(); i++ {
		name := at.Field(i).Name
		bf := bv.FieldByName(name)

		if bf.IsValid() && av.Field(i).Type().Name() == "string" && av.Field(i).Interface().(string) != "" {
			if bf.Type().Name() == "string" {
				bf.Set(av.Field(i))
			} else if bf.Type().Name() == "int" {
				newVal, _ := strconv.ParseInt(av.Field(i).Interface().(string), 10, 0)
				bf.SetInt(newVal)
			}

		} else if bf.IsValid() && av.Field(i).Type().Name() == "int" && bf.Type().Name() == "int" && bf.Interface().(int) != av.Field(i).Interface().(int) {
			bf.Set(av.Field(i))

		} else if bf.IsValid() && av.Field(i).Type().Name() == "uint" && bf.Type().Name() == "uint" && bf.Interface().(uint) != av.Field(i).Interface().(uint) {
			bf.Set(av.Field(i))

		}
	}

}

func EqualHead(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if strings.ToLower(strings.TrimSpace(v)) != strings.ToLower(strings.TrimSpace(b[i])) {
			return false
		}
	}
	return true
}

func AfterCreateUpdateOrderProgress(orderId uint, quantity int) {
	var order models.Order
	var product models.Product
	d := db.GetDB()

	d.Where("ID = ?", orderId).First(&order)
	d.Where("ID = ?", order.ProductID).First(&product)

	product.Quantity += quantity
	d.Save(&product)
}

func BeforeUpdateOrderProgress(id, orderId uint, quantity int) {
	var order models.Order
	var orderProgress models.OrderProgress
	var product models.Product

	d := db.GetDB()

	d.Where("ID = ?", id).First(&orderProgress)
	d.Where("ID = ?", orderId).First(&order)
	d.Where("ID = ?", order.ProductID).First(&product)

	product.Quantity -= orderProgress.QuantityReceived
	d.Save(&product)
}

func AfterCreateProductOut(productId uint, quantity int) {
	var product models.Product

	d := db.GetDB()
	d.Where("ID = ?", productId).First(&product)

	product.Quantity -= quantity
	d.Save(&product)
}
