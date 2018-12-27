package server

import (
	"github.com/kmaguswira/salestock-api/config"
)

func Init() {
	config := config.GetConfig()
	r := SetupRouter(config.GetString("name"))

	r.Run(config.GetString("server.port"))
}
