package main

import (
	"github.com/gin-gonic/gin"
	"github.com/importer/src/routes"
)

func main() {

	app := gin.New()
	routes.SetupRouter(app)

	app.Run(":4000")
}
