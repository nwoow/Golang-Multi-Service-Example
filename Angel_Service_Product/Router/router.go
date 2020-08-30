package Router

import (
	controllers "Angel/Angel_Service_Product/Controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	// r := gin.New()
	r := gin.Default()
	// r.Use(Logger.Logger(logger), gin.Recovery())
	r.POST("/addproduct", controllers.Addproduct)
	r.GET("/listproducts", controllers.Getproducts) // new
	r.Run(":8080")
}
