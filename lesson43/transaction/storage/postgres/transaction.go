package postgres

import (
	"database/sql"
	"transaction/models"
	"github.com/google/uuid"
)

type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{Db: db}
}

func (r *TransactionRepo) Create(transaction *models.Transaction) error {
	_, err := r.Db.Exec("INSERT INTO transactions (id, card_id, amount, terminal_id, transaction_type) VALUES ($1, $2, $3, $4, $5)",
		uuid.NewString(), transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType)
	return err
}

func (r *TransactionRepo) GetById(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.Db.QueryRow("SELECT id, card_id, amount, terminal_id, transaction_type FROM transactions WHERE id = $1", id).
		Scan(&transaction.Id, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepo) Update(id string, transaction *models.Transaction) error {
	_, err := r.Db.Exec("UPDATE transactions SET card_id = $1, amount = $2, terminal_id = $3, transaction_type = $4 WHERE id = $5",
		transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType, id)
	return err
}

func (r *TransactionRepo) Delete(id string) error {
	_, err := r.Db.Exec("DELETE FROM transactions WHERE id = $1", id)
	return err
}
