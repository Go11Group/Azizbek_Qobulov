package postgres

import (
	pb "LibraryService/server/genproto/BookService"
	"context"
	"database/sql"
)

type PostgresServer struct {
	pb.UnimplementedLibraryServiceServer
	DB *sql.DB
}

func (s *PostgresServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	var bookID string
	err := s.DB.QueryRow(
		"INSERT INTO books (title, author, year_published) VALUES ($1, $2, $3) RETURNING book_id",
		req.Title, req.Author, req.YearPublished).Scan(&bookID)
	if err != nil {
		return nil, err
	}
	return &pb.AddBookResponse{BookId: bookID}, nil
}

func (s *PostgresServer) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	rows, err := s.DB.Query("SELECT book_id, title, author, year_published FROM books WHERE title = $1 OR author = $1", req.Query)
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

func (s *PostgresServer) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	res, err := s.DB.Exec("DELETE FROM books WHERE book_id = $1", req.BookId)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	return &pb.BorrowBookResponse{Success: rowsAffected > 0}, nil
}
