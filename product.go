package main

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateProductDB(db *gorm.DB, product *Product) {
	result := db.Create(product)
	if result.Error != nil {
		return
	}
	fmt.Println("Product created successfully")
}

func getAllProductsDB(db *gorm.DB) {
	var products []Product

	result := db.Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get all product success")
	fmt.Println(products)
}

func getProductByIDDB(db *gorm.DB, id int) {
	var products []Product

	result := db.First(&products, id)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get product by ID success")
	fmt.Println(products)
}

func getProductByNameDB(db *gorm.DB, name string) {
	var products []Product

	result := db.Where("name LIKE ?", "%"+name+"%").Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get product by name success")
	fmt.Println(products)
}

func getProductsByCategoryDB(db *gorm.DB, categoryID int) {
	var products []Product

	result := db.Where("category_id = ?", categoryID).Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get products by category success")
	fmt.Println(products)
}

func getProductsByPriceRangeDB(db *gorm.DB, minPrice float64, maxPrice float64) {
	var products []Product

	result := db.Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get products by price range success")
	fmt.Println(products)
}

func getProductsLowStockDB(db *gorm.DB, stock int) {
	var products []Product

	result := db.Where("quantity < ?", stock).Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get products low stock success")
	fmt.Println(products)
}

func getProductsByCategoryAndPriceDB(db *gorm.DB, categoryID int, minPrice float64, maxPrice float64) {
	var products []Product

	result := db.Where("category_id = ? AND price BETWEEN ? AND ?", categoryID, minPrice, maxPrice).Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get products by category and price success")
	fmt.Println(products)
}

func getProductsSortedByPriceDB(db *gorm.DB, order string) {
	var products []Product

	result := db.Order("price " + order).Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get products sorted by price success")
	fmt.Println(products)
}

func getTopExpensiveProductsDB(db *gorm.DB, limit int) {
	var products []Product

	result := db.Order("price DESC").Limit(limit).Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get top expensive products success")
	fmt.Println(products)
}

func updateProductByIDDB(db *gorm.DB, id int, updatedProduct *Product) {
	if updatedProduct.Price < 0 || updatedProduct.Quantity < 0 {
		panic("Price and stock cannot be negative")
	}

	var product Product

	result := db.First(&product, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Model(&product).Updates(updatedProduct)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product by ID success")
	fmt.Println(product)
}

func updateProductPriceDB(db *gorm.DB, id int, newPrice float64) {
	if newPrice < 0 {
		panic("Price cannot be negative")
	}

	var product Product

	result := db.First(&product, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Model(&product).Update("price", newPrice)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product price success")
	fmt.Println(product)
}

func updateProductStockDB(db *gorm.DB, id int, change int) {
	var product Product

	result := db.First(&product, id)

	if result.Error != nil {
		panic(result.Error)
	}

	newStock := product.Quantity + int32(change)
	if newStock < 0 {
		panic("Stock cannot be negative")
	}

	result = db.Model(&product).Update("quantity", newStock)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product stock success")
	fmt.Println(product)
}

func softDeleteProductDB(db *gorm.DB, id int) {
	var product Product

	result := db.First(&product, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Delete(&product)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Soft delete product success")
	fmt.Println(product)
}

func hardDeleteProductDB(db *gorm.DB, id int) {
	var orderCount int64

	db.Model(&Order{}).Where("product_id = ?", id).Count(&orderCount)

	if orderCount > 0 {
		panic("Cannot hard delete product: product has order history")
	}

	var product Product

	result := db.Unscoped().First(&product, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Unscoped().Delete(&product)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Hard delete product success")
	fmt.Println(product)
}

func getDeletedProductsDB(db *gorm.DB) {
	var products []Product

	result := db.Unscoped().Where("deleted_at IS NOT NULL").Find(&products)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get deleted products success")
	fmt.Println(products)
}

func restoreProductDB(db *gorm.DB, id int) {
	var product Product

	result := db.Unscoped().First(&product, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Unscoped().Model(&product).Update("deleted_at", nil)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Restore product success")
	fmt.Println(product)
}


