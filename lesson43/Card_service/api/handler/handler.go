package handler

import (
	"database/sql"
	"Card_service/storage/postgres"
)

type handler struct {
	Card *postgres.CardRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		Card: postgres.NewCardRepo(db),
	}
}
