package main

import (
	models "Angel/Angel_Service_Cart/Models"
	"Angel/Angel_Service_Cart/Router"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	models.ConnectDataBase() // new
	Router.Router()
}
