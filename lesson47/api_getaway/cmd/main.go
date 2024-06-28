package main

import (
	"api/api"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	router := api.NewRouter(conn)
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
