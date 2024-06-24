package postgres

import (
	"database/sql"
	"user_service/models"
	"github.com/google/uuid"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Create(user *models.User) error {
	_, err := u.Db.Exec("INSERT INTO users(id, name, phone) VALUES ($1, $2, $3)", uuid.NewString(), user.Name, user.Phone)
	return err
}

func (u *UserRepo) GetById(id string) (*models.User, error) {
	var user = models.User{Id: id}
	err := u.Db.QueryRow("SELECT name, phone FROM users WHERE id = $1", id).Scan(&user.Name, &user.Phone)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) Update(id string, user *models.User) error {
	_, err := u.Db.Exec("UPDATE users SET name = $1, phone = $2 WHERE id = $3", user.Name, user.Phone, id)
	return err
}

func (u *UserRepo) Delete(id string) error {
	_, err := u.Db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
