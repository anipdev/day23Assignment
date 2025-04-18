package models

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

type Inventory struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Location  string `json:"location"`
}

type Order struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Datetime  string `json:"datetime"`
}
