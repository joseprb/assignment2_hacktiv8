package Config

import (
	"assignment2_hacktiv8/Structs"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	Conn, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	Conn.AutoMigrate(Structs.Order{})
	Conn.AutoMigrate(Structs.Item{})
	return Conn
}
