package db

import (
	"fmt"
	"log"
	"reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kmaguswira/salestock-api/config"
	"github.com/kmaguswira/salestock-api/models"
)

var db *gorm.DB
var err error
var tables []interface{} = []interface{}{
	models.Product{},
}

func Init() {
	config := config.GetConfig()
	shadow, err := gorm.Open("sqlite3", config.GetString("db.core.path"))	

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}

	log.Println("Database connected")

	db = shadow
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

func Migrate() {
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
