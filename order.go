package main

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateOrderDB(db *gorm.DB, order *Order) {
	result := db.Create(order)
	if result.Error != nil {
		return
	}
	fmt.Println("Order created successfully")
}

func getOrdersByCustomerDB(db *gorm.DB, customerID int) {
	var orders []Order

	result := db.Where("customer_id = ?", customerID).Find(&orders)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get orders by customer success")
	fmt.Println(orders)
}

func getOrdersByStatusDB(db *gorm.DB, status string) {
	var orders []Order

	result := db.Where("status = ?", status).Find(&orders)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get orders by status success")
	fmt.Println(orders)
}

func getOrdersByDateRangeDB(db *gorm.DB, startDate string, endDate string) {
	var orders []Order

	result := db.Where("DATE(created_at) BETWEEN ? AND ?", startDate, endDate).Find(&orders)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get orders by date range success")
	fmt.Println(orders)
}

type OrderDetail struct {
	ID           uint
	CustomerID   uint
	CustomerName string
	Email        string
	ProductID    uint
	ProductName  string
	Quantity     int32
	TotalPrice   float64
	Status       string
}

func getOrdersWithDetailsDB(db *gorm.DB) {
	var orderDetails []OrderDetail

	result := db.Table("orders").
		Select("orders.id, orders.customer_id, customers.name as customer_name, customers.email, orders.product_id, products.name as product_name, orders.quantity, orders.total_price, orders.status").
		Joins("JOIN customers ON customers.id = orders.customer_id").
		Joins("JOIN products ON products.id = orders.product_id").
		Scan(&orderDetails)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get orders with details success")
	fmt.Println(orderDetails)
}

func updateOrderStatusDB(db *gorm.DB, id int, status string) {
	if status != "Pending" && status != "Completed" && status != "Cancelled" {
		panic("Status must be Pending, Completed or Cancelled")
	}

	var order Order

	result := db.First(&order, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Model(&order).Update("status", status)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update order status success")
	fmt.Println(order)
}

