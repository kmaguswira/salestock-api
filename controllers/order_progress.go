package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/forms"
	"github.com/kmaguswira/salestock-api/models"
	"github.com/kmaguswira/salestock-api/utils"
)

type OrderProgressController struct{}

func (o OrderProgressController) FindOne(c *gin.Context) {
	var orderProgress models.OrderProgress
	db := db.GetDB()
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&orderProgress).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Model(orderProgress).Related(&orderProgress.Order)

	c.JSON(http.StatusOK, &orderProgress)
}

func (o OrderProgressController) Find(c *gin.Context) {
	var orderProgresss []models.OrderProgress
	db := db.QueryBuilder(c)
	db.Find(&orderProgresss)

	for i := range orderProgresss {
		db.Model(orderProgresss[i]).Related(&orderProgresss[i].Order)
	}

	c.JSON(http.StatusOK, &orderProgresss)
}

func (o OrderProgressController) Create(c *gin.Context) {
	var form forms.CreateOrderProgress
	var orderProgress models.OrderProgress
	var db = db.GetDB()

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	utils.CopyValue(form, &orderProgress)

	db.Create(&orderProgress)
	utils.AfterCreateUpdateOrderProgress(orderProgress.OrderID, orderProgress.QuantityReceived)

	db.Model(orderProgress).Related(&orderProgress.Order)

	c.JSON(http.StatusOK, &orderProgress)
}

func (o OrderProgressController) Update(c *gin.Context) {
	id := c.Param("id")
	var orderProgress models.OrderProgress
	var form forms.UpdateOrderProgress

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&orderProgress).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	utils.CopyValue(form, &orderProgress)

	utils.BeforeUpdateOrderProgress(orderProgress.ID, orderProgress.OrderID, orderProgress.QuantityReceived)
	db.Save(&orderProgress)
	utils.AfterCreateUpdateOrderProgress(orderProgress.OrderID, orderProgress.QuantityReceived)

	db.Model(orderProgress).Related(&orderProgress.Order)

	c.JSON(http.StatusOK, &orderProgress)
}

func (o OrderProgressController) Delete(c *gin.Context) {
	id := c.Param("id")
	var orderProgress models.OrderProgress
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&orderProgress).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&orderProgress)

	db.Model(orderProgress).Related(&orderProgress.Order)

	c.JSON(http.StatusOK, &orderProgress)

}
