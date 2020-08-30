package models

type Cart struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	Quantity   int    `json:"quantity"`
	Productid  int    `json:"productid"`
	Totalprice int    `json:"totlaprice"`
	User       int    `json:"user"`
}

type Availablequantity struct {
	ID        uint `json:"id" gorm:"primary_key"`
	Quantity  int  `json:"quantity"`
	Productid int  `json:"productid"`
}
