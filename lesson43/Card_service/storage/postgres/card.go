package postgres

import (
	"database/sql"
	"Card_service/models"
	"github.com/google/uuid"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

func (c *CardRepo) Create(card *models.Card) error {
	_, err := c.Db.Exec("INSERT INTO cards(id, number, user_id, level_id) VALUES ($1, $2, $3, $4)", uuid.NewString(), card.Number, card.UserId, card.LevelId)
	return err
}

func (c *CardRepo) GetById(id string) (*models.Card, error) {
	var card = models.Card{Id: id}
	err := c.Db.QueryRow("SELECT number, user_id, level_id FROM cards WHERE id = $1", id).Scan(&card.Number, &card.UserId, &card.LevelId)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (c *CardRepo) Update(id string, card *models.Card) error {
	_, err := c.Db.Exec("UPDATE cards SET number = $1, user_id = $2, level_id = $3 WHERE id = $4", card.Number, card.UserId, card.LevelId, id)
	return err
}

func (c *CardRepo) Delete(id string) error {
	_, err := c.Db.Exec("DELETE FROM cards WHERE id = $1", id)
	return err
}
