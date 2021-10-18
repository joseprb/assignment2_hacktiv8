package Structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string
	OrderedAt    time.Time
}
