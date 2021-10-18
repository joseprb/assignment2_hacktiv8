package main

import (
	"assignment2_hacktiv8/Config"
	"assignment2_hacktiv8/Controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := Config.InitDB()
	DBConn := &Controllers.DBConn{DB: db}

	router := gin.Default()

	router.GET("/orders", DBConn.GetOrders)
	router.GET("/orders/:id", DBConn.GetOrder)

	router.POST("/orders", DBConn.CreateOrder)

	router.PUT("/orders/:id", DBConn.UpdateOrder)

	router.DELETE("/orders/:id", DBConn.DeleteOrder)

	router.Run(":8080")
}
