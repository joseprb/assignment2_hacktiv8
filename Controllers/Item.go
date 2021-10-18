package Controllers

import (
	"assignment2_hacktiv8/Structs"
)

func (Conn *DBConn) CreateItem(i Structs.Item) {
	var item Structs.Item

	item.ItemCode = i.ItemCode
	item.Description = i.Description
	item.Quantity = i.Quantity
	item.OrderId = i.OrderId

	Conn.DB.Create(&item)
}
