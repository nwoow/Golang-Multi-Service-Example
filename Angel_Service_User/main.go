package main

import (
	models "Angel/Angel_Service_User/Models"
	"Angel/Angel_Service_User/Router"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	models.ConnectDataBase() // new
	Router.Router()
}
