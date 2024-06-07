package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Enter Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	name = strings.TrimSpace(name)

	_, err = conn.Write([]byte(name + "\n"))
	if err != nil {
		fmt.Println("Error sending name:", err)
		return
	}

	go func() {
		serverReader := bufio.NewReader(conn)
		for {
			message, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from server:", err)
				return
			}
			fmt.Print(message)
		}
	}()

	for {
		// fmt.Print("Enter message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		message = strings.TrimSpace(message)
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}
