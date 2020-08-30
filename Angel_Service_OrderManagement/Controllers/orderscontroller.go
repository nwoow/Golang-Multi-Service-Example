// controllers/books.go

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	models "Angel/Angel_Service_OrderManagement/Models"

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
	resp, err := http.Get("http://172.20.0.3:8080/listuser")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	in := []byte(body)
	var raw map[string]interface{}
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	raw["count"] = 1
	c.JSON(http.StatusOK, gin.H{"data": raw["data"]})
}
