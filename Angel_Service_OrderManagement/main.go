package main

import (
	models "Angel/Angel_Service_OrderManagement/Models"
	"Angel/Angel_Service_OrderManagement/Router"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	models.ConnectDataBase() // new
	Router.Router()
}
