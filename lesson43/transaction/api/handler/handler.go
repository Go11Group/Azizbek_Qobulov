package handler

import (
	"database/sql"
	"transaction/storage/postgres"
)

type handler struct {
	Transaction *postgres.TransactionRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		Transaction: postgres.NewTransactionRepo(db)}
}
