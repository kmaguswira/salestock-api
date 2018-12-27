package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/salestock-api/controllers"
	"github.com/kmaguswira/salestock-api/middlewares"
)

func SetupRouter(env string) *gin.Engine {
	var router *gin.Engine

	if env == "test" {
		router = gin.Default()
		gin.SetMode(gin.TestMode)
	} else {
		router = gin.New()

		router.Use(gin.Logger())
		router.Use(gin.Recovery())
		router.Static("/public", "./public")
	}

	router.Use(middlewares.CorsMiddleware())

	RouterV1(router)

	return router
}

func RouterV1(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		productGroup := v1.Group("product")
		{
			product := new(controllers.ProductController)
			productGroup.POST("/create", product.Create)
			productGroup.GET("/all", product.Find)
			productGroup.GET("/get/:id", product.FindOne)
			productGroup.PUT("/update/:id", product.Update)
			productGroup.DELETE("/delete/:id", product.Delete)
		}

		orderGroup := v1.Group("order")
		{
			order := new(controllers.OrderController)
			orderGroup.POST("/create", order.Create)
			orderGroup.GET("/all", order.Find)
			orderGroup.GET("/get/:id", order.FindOne)
			orderGroup.PUT("/update/:id", order.Update)
			orderGroup.DELETE("/delete/:id", order.Delete)
		}

		orderProgressGroup := v1.Group("order-progress")
		{
			orderProgress := new(controllers.OrderProgressController)
			orderProgressGroup.POST("/create", orderProgress.Create)
			orderProgressGroup.GET("/all", orderProgress.Find)
			orderProgressGroup.GET("/get/:id", orderProgress.FindOne)
			orderProgressGroup.PUT("/update/:id", orderProgress.Update)
			orderProgressGroup.DELETE("/delete/:id", orderProgress.Delete)
		}
	}

}
