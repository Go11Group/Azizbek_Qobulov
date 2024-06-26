package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"LibraryService/Storage/postgres"
	pb "LibraryService/server/genproto/BookService"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
	db *sql.DB
}

func (s *server) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	var bookID string
	err := s.db.QueryRow(
		"INSERT INTO books (title, author, year_published) VALUES ($1, $2, $3) RETURNING book_id",
		req.Title, req.Author, req.YearPublished).Scan(&bookID)
	if err != nil {
		return nil, err
	}
	return &pb.AddBookResponse{BookId: bookID}, nil
}

func (s *server) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	rows, err := s.db.Query("SELECT book_id, title, author, year_published FROM books WHERE title = $1 OR author = $1", req.Query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*pb.Book
	for rows.Next() {
		var book pb.Book
		if err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublished); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return &pb.SearchBookResponse{Books: books}, nil
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	res, err := s.db.Exec("DELETE FROM books WHERE book_id = $1", req.BookId)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &pb.BorrowBookResponse{Success: rowsAffected > 0}, nil
}

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLibraryServiceServer(s, &server{db: db})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
