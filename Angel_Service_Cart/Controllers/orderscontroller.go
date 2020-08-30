// controllers/books.go

package controllers

import (
	"net/http"

	models "Angel/Angel_Service_Cart/Models"

	"github.com/gin-gonic/gin"
)

func Addtocart(c *gin.Context) {
	// Validate input
	var input models.Cart
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addtocart := models.Cart{Title: input.Title, Quantity: input.Quantity, Productid: input.Productid, User: input.User}
	models.DB.Create(&addtocart)

	c.JSON(http.StatusOK, gin.H{"data": addtocart})
}
