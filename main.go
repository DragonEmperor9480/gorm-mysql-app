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

	//Insert data
	users := []User{
		{Name: "Alice", Email: "test1234@gmail.com", Age: 25},
		{Name: "Alice2", Email: "test12341@gmail.com", Age: 26},
		{Name: "Alice3", Email: "test12342@gmail.com", Age: 27},
	}

	result := db.Create(&users)

	if result.Error != nil {
		fmt.Println("Error inserting Data:", result.Error)
	} else {
		fmt.Println("Data inserted sucessfully!")
	}

	//Fetch data
	var FetchUsers []User
	res := db.Find((&FetchUsers))
	if res.Error != nil {
		fmt.Println("Error fetching users:", result.Error)
	} else {
		fmt.Println("Fetched users:")
		for _, user := range FetchUsers {
			fmt.Printf("ID: %d, Name: %s,Age:%d\n ", user.ID, user.Name, user.Age)
		}
	}
}
