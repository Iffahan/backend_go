package main

import (
	"fmt"
	db "go-fiber-test/database"
	r "go-fiber-test/routes"

	m "go-fiber-test/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	db.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	db.DBConn.AutoMigrate(&m.Dogs{})
}

func main() {
	app := fiber.New()
	initDatabase()
	r.InetRoutes(app)
	app.Listen(":3000")
}
