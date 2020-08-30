package models

type Product struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Details  string `json:"details"`
}
