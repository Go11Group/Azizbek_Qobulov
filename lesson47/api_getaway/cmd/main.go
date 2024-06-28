package main

import (
	"api/api"
	"log"

	"google.golang.org/api/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	r := api.NewRouter(conn)
	r.Run(":8080")
}
