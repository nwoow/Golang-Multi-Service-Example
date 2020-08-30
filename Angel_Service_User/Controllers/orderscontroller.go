// controllers/books.go

package controllers

import (
	"net/http"

	models "Angel/Angel_Service_User/Models"

	"github.com/gin-gonic/gin"
)

func Addproduct(c *gin.Context) {
	// Validate input
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addproduct := models.Product{Title: input.Title, Quantity: input.Quantity, Details: input.Details}
	models.DB.Create(&addproduct)

	c.JSON(http.StatusOK, gin.H{"data": addproduct})
}

func Getproducts(c *gin.Context) {
	var Product []models.Product
	models.DB.Find(&Product)

	c.JSON(http.StatusOK, gin.H{"data": Product})
}

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

func Adduser(c *gin.Context) {
	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adduser := models.User{Name: input.Name, Email: input.Email, Password: input.Password}
	models.DB.Create(&adduser)

	c.JSON(http.StatusOK, gin.H{"data": adduser})
}

func Listuser(c *gin.Context) {
	var User []models.User
	models.DB.Find(&User)

	c.JSON(http.StatusOK, gin.H{"data": User})
}
