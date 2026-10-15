package main

import (
	"fmt"

	"gorm.io/gorm"
)

func CreateCategoryDB(db *gorm.DB, category *Category) {
	result := db.Create(category)
	if result.Error != nil {
		return
	}
	fmt.Println("Category created successfully")
}

func deleteCategoryByIDDB(db *gorm.DB, id int) {
	var productCount int64

	db.Model(&Product{}).Where("category_id = ?", id).Count(&productCount)

	if productCount > 0 {
		panic("Cannot delete category: category has existing products")
	}

	var category Category

	result := db.First(&category, id)

	if result.Error != nil {
		panic(result.Error)
	}

	result = db.Delete(&category)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Delete category by ID success")
	fmt.Println(category)
}

