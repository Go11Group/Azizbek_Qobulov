package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

type Product struct {
	ID            int
	Name          string
	Description   string
	Price         float64
	StockQuantity int
}

func main() {
	conn := "postgres://postgres:root@localhost:5432/master?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// users := []User{
	//     {Username: "user1", Email: "user1@example.com", Password: "password1"},
	//     {Username: "user2", Email: "user2@example.com", Password: "password2"},
	//     {Username: "user3", Email: "user3@example.com", Password: "password3"},
	//     {Username: "user4", Email: "user4@example.com", Password: "password4"},
	//     {Username: "user5", Email: "user5@example.com", Password: "password5"},
	//     {Username: "user6", Email: "user6@example.com", Password: "password6"},
	//     {Username: "user7", Email: "user7@example.com", Password: "password7"},
	//     {Username: "user8", Email: "user8@example.com", Password: "password8"},
	//     {Username: "user9", Email: "user9@example.com", Password: "password9"},
	//     {Username: "user10",Email: "user10@example.com", Password: "password10"},
	// }

	// for _, user := range users {
	//     _, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	// }

	// products := []Product{
	//     {Name: "Product1", Description: "Description1", Price: 10.99, StockQuantity: 100},
	//     {Name: "Product2", Description: "Description2", Price: 20.99, StockQuantity: 200},
	//     {Name: "Product3", Description: "Description3", Price: 30.99, StockQuantity: 300},
	//     {Name: "Product4", Description: "Description4", Price: 40.99, StockQuantity: 400},
	//     {Name: "Product5", Description: "Description5", Price: 50.99, StockQuantity: 500},
	//     {Name: "Product6", Description: "Description6", Price: 60.99, StockQuantity: 600},
	//     {Name: "Product7", Description: "Description7", Price: 70.99, StockQuantity: 700},
	//     {Name: "Product8", Description: "Description8", Price: 80.99, StockQuantity: 800},
	//     {Name: "Product9", Description: "Description9", Price: 90.99, StockQuantity: 900},
	//     {Name: "Product10",Description: "Description10",Price: 100.99,StockQuantity: 1000},
	// }

	// for _, product := range products {
	//     _, err := db.Exec("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4)", product.Name, product.Description, product.Price, product.StockQuantity)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	// }

	fmt.Println("Data inserted successfully")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	_, err = db.Query("SELECT * FROM Users")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("SELECT * FROM Products")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("UPDATE Users SET password = qwerty WHERE id = 4;")
	db.Exec("UPDATE Users SET Name = Azizbek WHERE name = user2;")
	db.Exec("UPDATE Products SET price = 77.9 WHERE id = 5;")
	db.Exec("UPDATE Products SET stockquantity = 350 WHERE stockquantity = 300;")

	fmt.Println("Data updated successfully")

	_, err = db.Query("SELECT * FROM Users")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("SELECT * FROM Products")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("DELETE FROM Users WHERE id = 3")
	db.Exec("DELETE FROM Products WHERE stockquantity < 400")
	fmt.Println("Data deleted successfully")

	_, err = db.Query("SELECT * FROM Users")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("SELECT * FROM Products")
	if err != nil {
		log.Fatal(err)
	}
}
