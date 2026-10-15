package main

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description string
	Products    []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Quantity    int32
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	Orders      []Order  `gorm:"foreignKey:ProductID"`
}

type Customer struct {
	gorm.Model
	Name   string
	Email  string
	Phone  string
	Orders []Order `gorm:"foreignKey:CustomerID"`
}

type Order struct {
	gorm.Model
	CustomerID uint
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	ProductID  uint
	Product    Product  `gorm:"foreignKey:ProductID"`
	Quantity   int32
	TotalPrice float64
	Status     string
}
