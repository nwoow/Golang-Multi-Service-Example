package Router

import (
	controllers "Angel/Angel_Service_Cart/Controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	// r := gin.New()
	r := gin.Default()
	// r.Use(Logger.Logger(logger), gin.Recovery())
	r.POST("/addtocart", controllers.Addtocart)
	r.Run(":8080")
}
