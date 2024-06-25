package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	fmt.Println("\nTarjima qilmoqchi bo'lgan so'zni yozing:")
	for {
		var translation string
		var word string
		fmt.Scanln(&word)

		client, err := rpc.DialHTTP("tcp", "localhost:8080")
		if err != nil {
			log.Fatal("Dialing error:", err)
		}

		err = client.Call("Translate.Translator", word, &translation)
		if err != nil {
			log.Fatal("Call error:", err)
		}

		fmt.Println("Tarjima:", translation)
	}
}
