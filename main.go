package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = ""
	dbname   = "store_db"
)

func main() {
	dbconn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(mysql.Open(dbconn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(db)
	fmt.Println("Database connected successfully")

	db.AutoMigrate(
		&Category{},
		&Product{},
		&Customer{},
		&Order{},
	)
	fmt.Println("Migration success")

	//Create

	// สร้างข้อมูลของ Category
	// newCategory := &Category{
	// 	Name:        "Electronics",
	// 	Description: "เครื่องใช้ไฟฟ้าประจำบ้าน",
	// }

	// CreateCategoryDB(db, newCategory)

	//สร้างข้อมูลของ Product

	// newProduct := &Product{
	// 	Name:        "Fans",
	// 	Description: "Hatari",
	// 	Price:       299,
	// 	Quantity:    10,
	// 	CategoryID:  5,
	// }

	// CreateProductDB(db, newProduct)

	//สร้างข้อมูลของ Customer

	// newCustomer := &Customer{
	// 	Name:  "Bannaruj Limsomwong",
	// 	Email: "s6652410030@sau.ac.th",
	// 	Phone: "0823080759",
	// }

	// CreateCustomerDB(db, newCustomer)

	//สร้างข้อมูลของ Order

	// newOrder := &Order{
	// 	CustomerID: 9,
	// 	ProductID:  5,
	// 	Quantity:   2,
	// 	TotalPrice: 598,
	// 	Status:     "Pending",
	// }

	// CreateOrderDB(db, newOrder)

	//READ

	//Product

	//ดูสินค้าทั้งหมด
	// getAllProductsDB(db)

	//ค้าหาสินค้าจากID
	// getProductByIDDB(db, 5)

	//ค้นหาจากชื่อบางส่วนด้วยLike
	// getProductByNameDB(db, "F")

	//ค้นหาตาม Category ID
	// getProductsByCategoryDB(db, 3)

	//ตาม CategoryID
	// getProductsByPriceRangeDB(db, 1000, 5000)

	//ค้นหาตามช่วงMin-Max
	// getProductsLowStockDB(db, 10)

	//ค้นหาสินค้าที่สต็อกน้อยกว่ากำหนด
	// getProductsByCategoryAndPriceDB(db, 3, 1000, 5000)

	//ค้นหาตามประเภทและช่วงราคา
	// getProductsSortedByPriceDB(db, "DESC")

	//เรียงน้อยสุด สูงสุด
	// getTopExpensiveProductsDB(db, 5)

	//Customer

	//ค้นหาลูกค้าด้วย Email
	// getCustomerByEmailDB(db, "s6652410030@sau.ac.th")

	//Order

	//ดูรายการ Order ทั้งหมด
	// getOrdersByCustomerDB(db, 5)

	//ดูรายการ Order ที่มีสถานะ Pending
	// getOrdersByStatusDB(db, "Pending")

	// getOrdersByDateRangeDB(db, "2022-01-01", "2026-12-31")

	// getOrdersWithDetailsDB(db)

	//UPDATE

	//Product

	// updateProductByIDDB(db, 5, &Product{Price: 399})

	// updateProductPriceDB(db, 5, 299)

	// updateProductStockDB(db, 5, 10)

	//Customer

	// updateCustomerByIDDB(db, 5, &Customer{Name: "Bannaruj Limsomwong", Email: "[EMAIL_ADDRESS]", Phone: "0823080759"})

	//Order

	// updateOrderStatusDB(db, 5, "Completed")

	//Delete

	//Customer

	// deleteCustomerByIDDB(db, 5)

	//Product

	// deleteProductByIDDB(db, 5)

	// softDeleteProductDB(db, 5)

	// hardDeleteProductDB(db, 5)

	// getDeletedProductsDB(db)

	// restoreProductDB(db, 5)
}
