package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/importer/src/controllers"
)

func SetupRouter(app *gin.Engine) {

	routes := app.Group("/importer")
	routes.POST("/create", controllers.CreateProducts)
}
