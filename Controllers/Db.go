package Controllers

import "github.com/jinzhu/gorm"

type DBConn struct {
	DB *gorm.DB
}
