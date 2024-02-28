package main

import (
	"os"

	"github.com/devanfer02/nosudes-be/bootstrap/database/mysql"
	"github.com/devanfer02/nosudes-be/bootstrap/env"
	"github.com/devanfer02/nosudes-be/middleware"
	"github.com/devanfer02/nosudes-be/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	mysqldb := mysql.NewMysqlConn()
	defer mysqldb.Close()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "down":
			mysql.DropAllTables(mysqldb)
			mysql.MigrateUp(mysqldb)
		case "seeders":
			mysql.GenerateSeeders(mysqldb)
		}
	}

	if env.ProcEnv.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()
	app.Use(middleware.CORS())

	routes.InitRouter(app, mysqldb)

	app.Run(env.ProcEnv.ServerAddress)
}
