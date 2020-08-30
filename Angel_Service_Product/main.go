package main

import (
	models "Angel/Angel_Service_Product/Models"
	"Angel/Angel_Service_Product/Router"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	models.ConnectDataBase() // new
	Router.Router()
}
