package main

import (
	"context"
	"log"
	"time"

	pb "LibraryService/server/genproto/BookService"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLibraryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// AddBook
	addBookRes, err := c.AddBook(ctx, &pb.AddBookRequest{
		Title:         "Sample Book",
		Author:        "John Doe",
		YearPublished: 2022,
	})
	if err != nil {
		log.Fatalf("could not add book: %v", err)
	}
	log.Printf("Added Book ID: %s", addBookRes.BookId)

	// SearchBook
	searchBookRes, err := c.SearchBook(ctx, &pb.SearchBookRequest{Query: "Sample Book"})
	if err != nil {
		log.Fatalf("could not search book: %v", err)
	}
	log.Printf("Found Books: %v", searchBookRes.Books)

	// BorrowBook
	borrowBookRes, err := c.BorrowBook(ctx, &pb.BorrowBookRequest{
		BookId: addBookRes.BookId,
		UserId: "user123",
	})
	if err != nil {
		log.Fatalf("could not borrow book: %v", err)
	}
	log.Printf("Borrow Book Success: %v", borrowBookRes.Success)
}
