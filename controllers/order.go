package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/forms"
	"github.com/kmaguswira/salestock-api/models"
	"github.com/kmaguswira/salestock-api/utils"
)

type OrderController struct{}

func (o OrderController) FindOne(c *gin.Context) {
	var order models.Order
	db := db.GetDB()
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Model(order).Related(&order.Product)
	db.Model(order).Related(&order.OrderProgress)

	c.JSON(http.StatusOK, &order)
}

func (o OrderController) Find(c *gin.Context) {
	var orders []models.Order
	db := db.QueryBuilder(c)
	db.Find(&orders)

	for i := range orders {
		db.Model(orders[i]).Related(&orders[i].Product)
		db.Model(orders[i]).Related(&orders[i].OrderProgress)
	}

	c.JSON(http.StatusOK, &orders)
}

func (o OrderController) Create(c *gin.Context) {
	var form forms.CreateOrder
	var order models.Order
	var db = db.GetDB()

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	utils.CopyValue(form, &order)

	db.Create(&order)

	db.Model(order).Related(&order.Product)

	c.JSON(http.StatusOK, &order)
}

func (o OrderController) Update(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	var form forms.UpdateOrder

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	utils.CopyValue(form, &order)

	db.Save(&order)

	db.Model(order).Related(&order.Product)

	c.JSON(http.StatusOK, &order)
}

func (o OrderController) Delete(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&order).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&order)

	db.Model(order).Related(&order.Product)

	c.JSON(http.StatusOK, &order)

}
