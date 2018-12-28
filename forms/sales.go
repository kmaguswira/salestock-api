package forms

import (
	"github.com/kmaguswira/salestock-api/models"
)

type CreateSales struct {
	Note     string              `json:"note" binding:"required"`
	Products []models.ProductOut `json:"products" binding:"required"`
}
