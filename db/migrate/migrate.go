package migrate

import (
	"fmt"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/models"
	"log"
	"reflect"
)

var tables []interface{} = []interface{}{
	models.Product{},
	models.Order{},
	models.OrderProgress{},
	models.Sales{},
	models.ProductOut{},
}

func Migrate() {
	db := db.GetDB()
	for _, model := range tables {

		if !db.HasTable(model) {
			db.CreateTable(model)
		} else {
			log.Println("Table", reflect.TypeOf(model).Name(), "already exists")
		}

		if err := db.AutoMigrate(model).Error; err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Auto migrating", reflect.TypeOf(model).Name(), "...")
		}
	}
}
