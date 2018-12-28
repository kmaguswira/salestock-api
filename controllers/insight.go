package controllers

import (
	"encoding/csv"
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/models"
)

type InsightController struct{}

func (i InsightController) ValueProductInsight(c *gin.Context) {
	data := ValueProduct()
	writeCsv("product", data)
	c.JSON(http.StatusOK, data)
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
				"datetime":   sales[i].CreatedAt.Format("2006-01-02 15:04:05"),
			}

			totalOmzet += sales[i].ProductOuts[j].TotalPrice
			totalProfit += sales[i].ProductOuts[j].TotalPrice - (mean * sales[i].ProductOuts[j].Quantity)
			totalQuantity += sales[i].ProductOuts[j].Quantity
			dataSales = append(dataSales, row)
		}
	}

	data := map[string]interface{}{
		"sales":         dataSales,
		"datePrinted":   time.Now().Format("02 January 2006"),
		"dateRange":     startDate.Format("02 January 2006") + " - " + endDate.Format("02 January 2006"),
		"totalOmzet":    totalOmzet,
		"totalProfit":   totalProfit,
		"totalSales":    totalSales,
		"totalQuantity": totalQuantity,
	}
	writeCsv("sales", data)
	c.JSON(http.StatusOK, data)
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

func writeCsv(typeCsv string, data map[string]interface{}) {
	file, _ := os.Create(filepath.Join("public", typeCsv+".csv"))
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if typeCsv == "product" {
		row1 := []string{"Printed", data["date"].(string)}
		row2 := []string{"Total SKU", strconv.Itoa(data["totalSku"].(int))}
		row3 := []string{"Total Quantity SKU", strconv.Itoa(data["totalQuantitySku"].(int))}
		row4 := []string{"Total Price SKU", strconv.Itoa(data["totalPriceSku"].(int))}
		row5 := []string{"SKU", "Name", "Quantity", "Mean of Buying Price", "Total"}
		writer.Write(row1)
		writer.Write(row2)
		writer.Write(row3)
		writer.Write(row4)
		writer.Write(row5)
	} else {
		row1 := []string{"Printed", data["datePrinted"].(string)}
		row2 := []string{"Date Range", data["dateRange"].(string)}
		row3 := []string{"Total Omzet", strconv.Itoa(data["totalOmzet"].(int))}
		row4 := []string{"Total Profit", strconv.Itoa(data["totalProfit"].(int))}
		row5 := []string{"Total Sales", strconv.Itoa(data["totalSales"].(int))}
		row6 := []string{"Total Quantity", strconv.Itoa(data["totalQuantity"].(int))}
		row7 := []string{"ID", "Order ID", "Date time", "SKU", "Name", "Quantity", "Sell Price", "Total", "Buy Price", "Profit"}

		writer.Write(row1)
		writer.Write(row2)
		writer.Write(row3)
		writer.Write(row4)
		writer.Write(row5)
		writer.Write(row6)
		writer.Write(row7)

	}

	if typeCsv == "product" {
		for _, value := range data["products"].([]interface{}) {
			row := []string{
				value.(map[string]interface{})["sku"].(string),
				value.(map[string]interface{})["name"].(string),
				strconv.Itoa(value.(map[string]interface{})["quantity"].(int)),
				strconv.Itoa(value.(map[string]interface{})["meanPrice"].(int)),
				strconv.Itoa(value.(map[string]interface{})["total"].(int)),
			}
			writer.Write(row)
		}
	} else {
		for _, value := range data["sales"].([]interface{}) {

			row := []string{
				strconv.Itoa(int(value.(map[string]interface{})["salesID"].(uint))),
				value.(map[string]interface{})["note"].(string),
				value.(map[string]interface{})["datetime"].(string),
				value.(map[string]interface{})["sku"].(string),
				value.(map[string]interface{})["name"].(string),
				strconv.Itoa(value.(map[string]interface{})["quantity"].(int)),
				strconv.Itoa(value.(map[string]interface{})["sellPrice"].(int)),
				strconv.Itoa(value.(map[string]interface{})["totalPrice"].(int)),
				strconv.Itoa(value.(map[string]interface{})["buyPrice"].(int)),
				strconv.Itoa(value.(map[string]interface{})["profit"].(int)),
			}
			writer.Write(row)
		}
	}

}
