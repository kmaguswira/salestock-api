package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/forms"
	"github.com/kmaguswira/salestock-api/models"
	"github.com/kmaguswira/salestock-api/utils"
)

type ProductController struct{}

func (p ProductController) FindOne(c *gin.Context) {
	var product models.Product
	db := db.GetDB()
	id := c.Param("id")

	db.Where("id = ?", id).First(&product)

	c.JSON(http.StatusOK, &product)
}

func (p ProductController) Find(c *gin.Context) {
	var products []models.Product
	db := db.QueryBuilder(c)
	db.Find(&products)

	c.JSON(http.StatusOK, &products)
}

func (p ProductController) Create(c *gin.Context) {
	var form forms.CreateProduct
	var product models.Product
	var db = db.GetDB()

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	utils.CopyValue(form, &product)

	db.Create(&product)

	c.JSON(http.StatusOK, &product)
}

func (p ProductController) Update(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	var form forms.UpdateProduct

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	utils.CopyValue(form, &product)

	db.Save(&product)

	c.JSON(http.StatusOK, &product)
}

func (p ProductController) Delete(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&product)
	c.JSON(http.StatusOK, &product)

}
