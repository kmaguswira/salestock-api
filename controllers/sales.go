package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/models"
)

type SalesController struct{}

func (s SalesController) FindOne(c *gin.Context) {
	var sales models.Sales
	db := db.GetDB()
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&sales).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Model(sales).Related(&sales.ProductOuts)

	c.JSON(http.StatusOK, &sales)
}

func (s SalesController) Find(c *gin.Context) {
	var saless []models.Sales
	db := db.QueryBuilder(c)
	db.Find(&saless)

	for i := range saless {
		db.Model(saless[i]).Related(&saless[i].ProductOuts)
	}

	c.JSON(http.StatusOK, &saless)
}

func (s SalesController) Create(c *gin.Context) {
	var sales models.Sales
	var db = db.GetDB()

	db.Create(&sales)

	db.Model(sales).Related(&sales.ProductOuts)

	c.JSON(http.StatusOK, &sales)
}

func (s SalesController) Delete(c *gin.Context) {
	id := c.Param("id")
	var sales models.Sales
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&sales).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&sales)

	db.Model(sales).Related(&sales.ProductOuts)

	c.JSON(http.StatusOK, &sales)

}
