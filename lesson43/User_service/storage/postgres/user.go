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
	_, err := u.Db.Exec("insert into users(id, name, phone,age) values ($1, $2, $3,$4)", uuid.NewString(), user.Name, user.Phone,user.Age)
	return err
}

func (u *UserRepo) GetById(id string) (*models.User, error) {
	var user = models.User{Id: id}
	err := u.Db.QueryRow("select name, phone , age from users where id = $1", id).Scan(&user.Name, &user.Phone,&user.Age)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
