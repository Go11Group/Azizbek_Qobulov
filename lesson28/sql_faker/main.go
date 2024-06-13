package main

import (
	// "fmt"
	"sql_mock/moduls"
)

func main() {
	db, err := moduls.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var data = moduls.Data_holder("MOCK_DATA _sql.json")
	for _, x := range data {
		_, err := db.Exec(`INSERT INTO users(id,first_name,last_name,email,gender,age) values($1,$2,$3,$4,$5,$6)`, x.ID, x.Fist_name, x.Last_name, x.Email, x.Gender, x.Age)
		if err != nil {
			panic(err)
		}
	}
	// fmt.Println(data)
}
