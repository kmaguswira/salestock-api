package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/forms"
	"github.com/kmaguswira/salestock-api/models"
	"github.com/kmaguswira/salestock-api/utils"
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
		for j := range saless[i].ProductOuts {
			db.Model(saless[i].ProductOuts[j]).Related(&saless[i].ProductOuts[j].Product)
		}
	}

	c.JSON(http.StatusOK, &saless)
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

func (s SalesController) NewSales(c *gin.Context) {
	var form forms.CreateSales
	var sales models.Sales
	var db = db.GetDB()

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sales.Note = form.Note
	db.Create(&sales)

	var res []interface{}
	for _, product := range form.Products {
		product.SalesID = sales.ID
		product.Type = "Sales"
		product.Note = form.Note
		product.TotalPrice = product.Quantity * product.SellPrice

		db.Create(&product)
		utils.AfterCreateProductOut(product.ProductID, product.Quantity)

		db.Model(&product).Related(&product.Product)
		res = append(res, product)
	}
	c.JSON(http.StatusOK, gin.H{
		"sales":    sales,
		"products": res,
	})

}
