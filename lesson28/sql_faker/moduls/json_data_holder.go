package moduls

import (
	"encoding/json"
	"fmt"
	"os"
)

func Data_holder(nameOfFile string) []Person {
	f, err := os.OpenFile(nameOfFile, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var holder []Person
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&holder)
	if err != nil {
		fmt.Println("_________________")
		panic(err)
	}
	return holder
}
