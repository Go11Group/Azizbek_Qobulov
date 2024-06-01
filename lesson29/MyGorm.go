package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "fmt"
)

type TestUser struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	dsn := "postgres://postgres:root@localhost:5432/master?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// users := []TestUser{
	// 	{Name: "Jinxua", Email: "jinhua1@example.com", Age: 18},
	// 	{Name: "Jinzhu2", Email: "jinzhu2@example.com", Age: 19},
	// 	{Name: "Jinzhu3", Email: "jinzhu3@example.com", Age: 20},
	// 	{Name: "Jinzhu4", Email: "jinzhu4@example.com", Age: 21},
	// 	{Name: "Jinzhu5", Email: "jinzhu5@example.com", Age: 22},
	// 	{Name: "Jinzhu6", Email: "jinzhu6@example.com", Age: 23},
	// 	{Name: "Jinzhu7", Email: "jinzhu7@example.com", Age: 24},
	// 	{Name: "Jinzhu8", Email: "jinzhu8@example.com", Age: 25},
	// 	{Name: "Jinzhu9", Email: "jinzhu9@example.com", Age: 26},
	// 	{Name: "Jinzhu10", Email: "jinzhu10@example.com", Age: 27},
	// }

	// for _, user := range users {
	// 	result := db.Create(&user)
	// 	if result.Error != nil {
	// 		log.Fatal("failed to create user:", result.Error)
	// 	}
	// }
	// var user TestUser
	// db.First(&user,"age = ?",18)
	// fmt.Println(user)

	// db.Model(&user).Where(&user,"id = ?",3).Update("name","Jin Woo")
	// fmt.Println(user)

	// db.Delete(&user, "10")
}
