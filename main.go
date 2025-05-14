package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string `gorm:"unique"`
	Age   int
}

func main() {
	dsn := "root:0000@tcp(127.0.0.1:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connected successfully!", db)

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to auto-migrate schema!")
	}
	fmt.Println("✅ Migration complete!")

}
