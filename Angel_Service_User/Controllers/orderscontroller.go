// controllers/books.go

package controllers

import (
	"net/http"

	models "Angel/Angel_Service_User/Models"

	"github.com/gin-gonic/gin"
)

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
