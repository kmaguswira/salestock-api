package controllers

import (
	// "encoding/csv"
	"fmt"
	"math"
	"net/http"
	// "os"
	// "path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/models"
)

type InsightController struct{}

func (i InsightController) ValueProductInsight(c *gin.Context) {
	c.JSON(http.StatusOK, ValueProduct())
}

func (i InsightController) SalesInsight(c *gin.Context) {
	var start string
	var end string
	var isExist bool

	if start, isExist = c.GetQuery("start"); !isExist {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid query"))
		return
	}

	if end, isExist = c.GetQuery("end"); !isExist {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid query"))
		return
	}

	var sales []models.Sales
	var totalSales int
	totalOmzet := 0
	totalProfit := 0
	totalQuantity := 0
	db := db.GetDB()

	startDate, _ := time.Parse("2006-01-02", start)
	endDate, _ := time.Parse("2006-01-02", end)
	endDate = endDate.AddDate(0, 0, 1)

	db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Order("created_at asc").Find(&sales).Count(&totalSales)
	valueProduct := ValueProduct()
	fmt.Println(valueProduct["products"])

	var dataSales []interface{}
	for i := range sales {
		db.Model(sales[i]).Related(&sales[i].ProductOuts)

		for j := range sales[i].ProductOuts {
			db.Model(sales[i].ProductOuts[j]).Related(&sales[i].ProductOuts[j].Product)

			mean := searchMean(valueProduct["products"].([]interface{}), sales[i].ProductOuts[j].Product.Sku)

			row := map[string]interface{}{
				"salesID":    sales[i].ID,
				"note":       sales[i].Note,
				"sku":        sales[i].ProductOuts[j].Product.Sku,
				"name":       sales[i].ProductOuts[j].Product.Name,
				"quantity":   sales[i].ProductOuts[j].Quantity,
				"sellPrice":  sales[i].ProductOuts[j].SellPrice,
				"totalPrice": sales[i].ProductOuts[j].TotalPrice,
				"buyPrice":   mean,
				"profit":     sales[i].ProductOuts[j].TotalPrice - (mean * sales[i].ProductOuts[j].Quantity),
			}

			totalOmzet += sales[i].ProductOuts[j].TotalPrice
			totalProfit += sales[i].ProductOuts[j].TotalPrice - (mean * sales[i].ProductOuts[j].Quantity)
			totalQuantity += sales[i].ProductOuts[j].Quantity
			dataSales = append(dataSales, row)
		}
	}

	c.JSON(http.StatusOK, &gin.H{
		"sales":         dataSales,
		"datePrinted":   time.Now().Format("02 January 2006"),
		"dateRange":     startDate.Format("02 January 2006") + " - " + endDate.Format("02 January 2006"),
		"totalOmzet":    totalOmzet,
		"totalProfit":   totalProfit,
		"totalSales":    totalSales,
		"totalQuantity": totalQuantity,
	})
}

func searchMean(data []interface{}, sku string) int {
	for _, product := range data {
		if product.(map[string]interface{})["sku"] == sku {
			return product.(map[string]interface{})["meanPrice"].(int)
		}
	}
	return 0
}

func ValueProduct() gin.H {
	var products []models.Product
	var totalSku int
	totalQuantitySku := 0
	totalPriceSku := 0
	db := db.GetDB()

	db.Find(&products).Order("ID asc").Count(&totalSku)

	var data []interface{}
	for _, product := range products {
		db.Model(product).Related(&product.Orders)

		totalPrice := 0
		totalQuantity := 0
		for _, order := range product.Orders {
			totalPrice += order.TotalPrice
			totalQuantity += order.OrderQuantity
		}

		totalQuantitySku += product.Quantity
		totalPriceSku += totalPrice
		mean := int(math.Ceil(float64(totalPrice) / float64(totalQuantity)))
		row := map[string]interface{}{
			"sku":       product.Sku,
			"name":      product.Name,
			"quantity":  product.Quantity,
			"meanPrice": mean,
			"total":     product.Quantity * mean,
		}

		data = append(data, row)

	}

	return gin.H{
		"date":             time.Now().Format("02 January 2006"),
		"totalSku":         totalSku,
		"totalQuantitySku": totalQuantitySku,
		"totalPriceSku":    totalPriceSku,
		"products":         data,
	}
}

// func writeCsv(typeCsv string, data map[string]interface{}) {
// 	file, _ := os.Create(filepath.Join("tmp", typeCsv))
//     defer file.Close()

//     writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	if typeCsv == "sales" {
// 		row := []string{"Printed", data["datePrinted"]}
// 	}

//     for _, value := range data {
// 		row = []string{}
//         if err := writer.Write(value); err != nil {
// 			fmt.Println("error")
// 		}
//     }
// }
