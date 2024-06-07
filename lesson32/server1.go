package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	connections = make(map[net.Conn]string)
	mutex       sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}
	defer listener.Close()
	fmt.Print("Enter message: ")
	fmt.Println("Server is listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		mutex.Lock()
		delete(connections, conn)
		mutex.Unlock()
		conn.Close()
	}()

	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	name = strings.TrimSpace(name)

	mutex.Lock()
	connections[conn] = name
	mutex.Unlock()

	fmt.Println("Client connected:", name)
	broadcastMessage(conn, fmt.Sprintf("%s has joined the chat\n", name))

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			broadcastMessage(conn, fmt.Sprintf("%s has left the chat\n", name))
			return
		}
		message = strings.TrimSpace(message)
		fmt.Printf("Received message from %s: %s\n", name, message)
		broadcastMessage(conn, fmt.Sprintf("%s: %s\n", name, message))
	}
}

func broadcastMessage(sender net.Conn, message string) {
	mutex.Lock()
	defer mutex.Unlock()
	for conn := range connections {
		if conn != sender {
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error broadcasting message:", err)
			}
		}
	}
}
