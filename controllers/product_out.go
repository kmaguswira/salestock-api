package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/forms"
	"github.com/kmaguswira/salestock-api/models"
	"github.com/kmaguswira/salestock-api/utils"
)

type ProductOutController struct{}

func (p ProductOutController) FindOne(c *gin.Context) {
	var productOut models.ProductOut
	db := db.GetDB()
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&productOut).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Model(productOut).Related(&productOut.Product)
	db.Model(productOut).Related(&productOut.Sales)

	c.JSON(http.StatusOK, &productOut)
}

func (p ProductOutController) Find(c *gin.Context) {
	var productOuts []models.ProductOut
	db := db.QueryBuilder(c)
	db.Find(&productOuts)

	for i := range productOuts {
		db.Model(productOuts[i]).Related(&productOuts[i].Product)
		db.Model(productOuts[i]).Related(&productOuts[i].Sales)
	}

	c.JSON(http.StatusOK, &productOuts)
}

func (p ProductOutController) Create(c *gin.Context) {
	var form forms.CreateProductOut
	var productOut models.ProductOut
	var db = db.GetDB()

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	utils.CopyValue(form, &productOut)

	db.Create(&productOut)

	db.Model(productOut).Related(&productOut.Product)
	db.Model(productOut).Related(&productOut.Sales)

	c.JSON(http.StatusOK, &productOut)
}

func (p ProductOutController) Update(c *gin.Context) {
	id := c.Param("id")
	var productOut models.ProductOut
	var form forms.UpdateProductOut

	if err := c.BindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&productOut).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	utils.CopyValue(form, &productOut)

	db.Save(&productOut)

	db.Model(productOut).Related(&productOut.Product)
	db.Model(productOut).Related(&productOut.Sales)

	c.JSON(http.StatusOK, &productOut)
}

func (p ProductOutController) Delete(c *gin.Context) {
	id := c.Param("id")
	var productOut models.ProductOut
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&productOut).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&productOut)

	db.Model(productOut).Related(&productOut.Product)
	db.Model(productOut).Related(&productOut.Sales)

	c.JSON(http.StatusOK, &productOut)
}
