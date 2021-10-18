package Controllers

import (
	"assignment2_hacktiv8/Structs"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (Conn *DBConn) CreateOrder(c *gin.Context) {
	var order Structs.Order
	var items []Structs.Item

	order.CustomerName = c.PostForm("customer_name")
	order.OrderedAt = time.Now()

	Conn.DB.Create(&order)

	itemsJson := []byte(c.PostForm("items"))
	err := json.Unmarshal(itemsJson, &items)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(items); i++ {
		items[i].OrderId = uint(order.ID)
		Conn.DB.Create(&items[i])
	}

	result := gin.H{
		"result": order,
	}

	c.JSON(http.StatusCreated, result)
}

func (Conn *DBConn) GetOrders(c *gin.Context) {
	var orders []Structs.Order
	var results gin.H

	Conn.DB.Find(&orders)

	if len(orders) <= 0 {
		results = gin.H{
			"results": nil,
		}
	} else {
		results = gin.H{
			"results": orders,
		}
	}

	c.JSON(http.StatusOK, results)
}

func (Conn *DBConn) GetOrder(c *gin.Context) {
	var order Structs.Order
	var result gin.H

	id := c.Param("id")
	err := Conn.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": order,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (Conn *DBConn) UpdateOrder(c *gin.Context) {
	var result gin.H
	var order Structs.Order

	id := c.Query("id")
	err := Conn.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	var newOrder Structs.Order
	newOrder.CustomerName = c.PostForm("customer_name")
	newOrder.OrderedAt = time.Now()

	err = Conn.DB.Model(&order).Updates(newOrder).Error

	// var items []Structs.Item
	var newItems []Structs.Item
	var ni Structs.Item
	itemsJson := []byte(c.PostForm("items"))
	err = json.Unmarshal(itemsJson, &newItems)
	// fmt.Println(len(newite))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < len(newItems); i++ {
		fmt.Println(Conn.DB.First(&ni, newItems[i].ID).Error)
		err = Conn.DB.Model(&ni).Updates(newItems[i]).Error
		if err != nil {
			result = gin.H{
				"result": "update failed",
			}
			c.JSON(http.StatusInternalServerError, result)
			return
		}
	}

	result = gin.H{
		"result": "update data success",
	}
	c.JSON(http.StatusOK, newItems)
}

func (Conn *DBConn) DeleteOrder(c *gin.Context) {
	var (
		order  Structs.Order
		item   Structs.Item
		result gin.H
	)

	id := c.Param("id")
	err := Conn.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	err = Conn.DB.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	Conn.DB.Where("order_id = ?", id).First(&item)
	Conn.DB.Delete(&item)

	result = gin.H{
		"result": "data deleted successfully",
	}

	c.JSON(http.StatusOK, result)
}
