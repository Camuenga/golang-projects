// dbconnect project dbconnect.go
package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Email string `gorm:"uniqueIndex;not null"`
}

func InitDB() {

	username := "yami"
	password := "-camuenga!-"
	host := "127.0.0.1"
	port := "3306"
	dbname := "my_database"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := DB.DB()
	sqlDB.SetMaxIdleConns(10)  // Maximum idle connections in the pool
	sqlDB.SetMaxOpenConns(100) // Maximum open connections allowed

	if err != nil {
		log.Fatalf("Failed to get underlying SQL driver: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connection successfully established.")

}

func main() {
	InitDB()
	err := DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to run schema auto-migration:%v", err)
	}
	fmt.Println("Database migration completed")
}
