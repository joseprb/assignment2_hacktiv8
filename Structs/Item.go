package Structs

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	ItemCode    string
	Description string
	Quantity    int
	OrderId     uint
	Order       *Order
}
