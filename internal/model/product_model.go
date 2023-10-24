package model

type Product struct {
	ID          int     `gorm:"type:int;primary_key"`
	ProductName string  `gorm:"type:varchar(50)"`
	Price       float64 `gorm:"type:float"`
	Quantity    int     `gorm:"type:int"`
}