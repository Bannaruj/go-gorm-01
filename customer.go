package main

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateCustomerDB(db *gorm.DB, customer *Customer) {
	result := db.Create(customer)
	if result.Error != nil {
		return
	}
	fmt.Println("Customer created successfully")
}

func getCustomerByEmailDB(db *gorm.DB, email string) {
	var customer Customer

	result := db.Where("email = ?", email).First(&customer)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get customer by email success")
	fmt.Println(customer)
}

func updateCustomerByIDDB(db *gorm.DB, id int, updatedCustomer *Customer) {
	var customer Customer

	result := db.First(&customer, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Model(&customer).Updates(updatedCustomer)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update customer by ID success")
	fmt.Println(customer)
}

func deleteCustomerByIDDB(db *gorm.DB, id int) {
	var orderCount int64

	db.Model(&Order{}).Where("customer_id = ?", id).Count(&orderCount)

	if orderCount > 0 {
		panic("Cannot delete customer: customer has order history")
	}

	var customer Customer

	result := db.First(&customer, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Delete(&customer)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Delete customer by ID success")
	fmt.Println(customer)
}


