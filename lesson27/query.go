package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "master"
	password = "root"
)

type Cars struct {
	id    string
	Name  string
	Year  int
	Brand string
}

func main() {

	con := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(`
    SELECT
	id,
	Name,
	Year,
	Brand
    FROM
	Cars
	`)
	if err != nil {
		log.Println(err)
	}

	cars := []Cars{}

	for rows.Next() {
		Cars := Cars{}
		rows.Scan(
			&Cars.id,
			&Cars.Name,
			&Cars.Year,
			&Cars.Brand,
		)
		cars = append(cars, Cars)
	}

	count := 0
	for _, v := range cars {
		fmt.Println(v)
		count++
	}
	fmt.Println(count)
}

func delete(id int, db *sql.DB) {
	_, err := db.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
}

func update(st *Cars, db *sql.DB) {
	_, err := db.Exec("UPDATE Cars SET name = $1, Year = $2 , brand = $3", st.Name, st.Year, st.Brand)
	if err != nil {
		panic(err)
	}
}
