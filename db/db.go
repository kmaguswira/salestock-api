package db

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kmaguswira/salestock-api/config"
	"github.com/kmaguswira/salestock-api/models"
)

var db *gorm.DB
var err error
var tables []interface{} = []interface{}{
	models.Product{},
	models.Order{},
	models.OrderProgress{},
	models.Sales{},
	models.ProductOut{},
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

func QueryBuilder(c *gin.Context) *gorm.DB {
	orderQuery := "created_at desc"
	var limitQuery int64 = 1000
	var offsetQuery int64 = 0
	var whereQuery interface{}
	var whereKey []string
	var whereValue []interface{}

	if order, isExist := c.GetQuery("order"); isExist {
		orderQuery = order
	}

	if limit, isExist := c.GetQuery("limit"); isExist {
		limitQuery, _ = strconv.ParseInt(limit, 10, 64)

	}

	if offset, isExist := c.GetQuery("offset"); isExist {
		offsetQuery, _ = strconv.ParseInt(offset, 10, 64)
	}

	if where, isExist := c.GetQuery("where"); isExist {

		json.Unmarshal([]byte(where), &whereQuery)
		for i, _ := range whereQuery.(map[string]interface{}) {
			whereValue = append(whereValue, whereQuery.(map[string]interface{})[i].(string))
			whereKey = append(whereKey, i+" ?")
		}
	}

	return db.Order(orderQuery).Limit(limitQuery).Offset(offsetQuery).Where(strings.Join(whereKey[:], " AND "), whereValue...)
}
