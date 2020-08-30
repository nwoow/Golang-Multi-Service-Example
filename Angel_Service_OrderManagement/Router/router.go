package Router

import (
	controllers "Angel/Angel_Service_OrderManagement/Controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	// r := gin.New()
	r := gin.Default()

	r.POST("/adduser", controllers.Adduser)
	r.GET("/listuser", controllers.Listuser) // new
	r.Run(":8080")
}
