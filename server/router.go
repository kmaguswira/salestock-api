package server

import (
	"github.com/gin-gonic/gin"
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
	

}
