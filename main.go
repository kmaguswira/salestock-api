package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kmaguswira/salestock-api/config"
	"github.com/kmaguswira/salestock-api/db"
	"github.com/kmaguswira/salestock-api/db/migrate"
	"github.com/kmaguswira/salestock-api/server"
)

func main() {
	env := flag.String("e", "development", "")

	flag.Usage = func() {
		fmt.Println("FLAG	DESCRIPTION		DEFAULT			OPTIONS")
		fmt.Println("-e		Environment		development		development/migrate")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*env)
	db.Init()
	bootstrap(*env)
}

func bootstrap(env string) {
	if env == "migrate" {
		migrate.Migrate()
		os.Exit(1)
	} else {
		server.Init()
	}
}
