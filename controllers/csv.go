package controllers

import (
	"net/http"

	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/models"
	"github.com/kmaguswira/salestock-api/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type CSVController struct{}

func (cs CSVController) ImportFrom(c *gin.Context) {
	typeCSV := c.PostForm("type")

	if typeCSV != "product" && typeCSV != "order" && typeCSV != "sales" {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid type"))
		return
	}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filepath.Join("tmp", filename)); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	f, _ := os.Open(filepath.Join("tmp", filename))
	defer f.Close()

	lines, _ := csv.NewReader(f).ReadAll()

	productHead := []string{"SKU", "Nama Item", "Jumlah Sekarang"}
	orderHead := []string{"Waktu ", "SKU", "Nama Barang", "Jumlah Pemesanan", "Jumlah Diterima", "Harga Beli ", "Total", "Nomer Kwitansi", "Catatan"}
	// salesHead := []string{"SKU", "Nama Item", "Jumlah Sekarang"}

	if typeCSV == "product" && utils.EqualHead(lines[0], productHead) {
		lines = lines[1:]
		for _, line := range lines {
			quantity, _ := strconv.Atoi(line[2])
			product := models.Product{
				Sku:      line[0],
				Name:     line[1],
				Quantity: quantity,
			}
			db := db.GetDB()
			db.Create(&product)
		}
	} else if typeCSV == "order" && utils.EqualHead(lines[0], orderHead) {
		lines = lines[1:]
		for _, line := range lines {
			db := db.GetDB()
			var product models.Product
			db.Where(models.Product{Sku: line[1]}).Attrs(models.Product{Name: line[2], Quantity: 0}).FirstOrCreate(&product)

			orderQuantity, _ := strconv.Atoi(line[3])
			basePrice, _ := strconv.Atoi(strings.Replace(strings.Replace(line[4], "Rp", "", -1), ",", "", -1))
			totalPrice, _ := strconv.Atoi(strings.Replace(strings.Replace(line[5], "Rp", "", -1), ",", "", -1))
			var invoice string

			if line[7] != "(Hilang)" {
				invoice = line[7]
			}

			order := models.Order{
				ProductID:     product.ID,
				OrderQuantity: orderQuantity,
				BasePrice:     basePrice,
				TotalPrice:    totalPrice,
				Status:        "Incomplete",
				Invoice:       invoice,
			}

			var t time.Time
			t, err = time.Parse("2006/01/02 15:04", line[0])
			if err != nil {
				t, _ = time.Parse("2006/1/02 15:04", line[0])
			}

			db.Create(&order)

			order.CreatedAt = t
			db.Save(&order)

			progress := strings.Split(line[8], "; ")

			for _, p := range progress {
				if p != "Masih Menunggu" {

				}
			}

		}
	} else {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid csv"))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, typeCSV, typeCSV))
	// c.JSON(http.StatusOK, &order)
}
