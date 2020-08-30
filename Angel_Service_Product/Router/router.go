package Router

import (
	"log"
	"time"

	controllers "Angel/Angel_Service_Product/Controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	// r := gin.New()
	r := gin.Default()
	// r.Use(Logger.Logger(logger), gin.Recovery())
	r.POST("/addproduct", controllers.Addproduct)
	r.POST("/addtocart", controllers.Addtocart)
	r.POST("/adduser", controllers.Adduser)

	r.GET("/listproducts", controllers.Getproducts) // new
	r.GET("/listuser", controllers.Listuser)        // new

	// r.GET("/logconfiglist", controllers.Logconfiglist) // nea
	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	// r.GET("/books/:id", controllers.FindBook)     // new
	// r.PATCH("/books/:id", controllers.UpdateBook) // new
	r.Run(":8080")
}
