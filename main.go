package main

import (
	"github.com/yudisaputra/assignment-bookandlink/core"
	"github.com/yudisaputra/assignment-bookandlink/database"
	"github.com/yudisaputra/assignment-bookandlink/routes"
)

func main() {
	// load env
	core.LoadEnv()

	// load database
	database.Mysql()

	// route
	routes.Route()
}
