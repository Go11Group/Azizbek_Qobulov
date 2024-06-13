package postgres

import (
	"ExamNT/modul"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (u *UsersRepo) GetByID(id string) (*modul.User, error) {
	query := "SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1"
	row := u.db.QueryRow(query, id)
	var user modul.User
	if err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersRepo) Delete(id string) error {
	query := "UPDATE users SET deleted_at = $1 WHERE user_id = $2"
	_, err := u.db.Exec(query, time.Now().Unix(), id)
	return err
}

func (u *UsersRepo) Update(user *modul.User) error {
	query := "UPDATE users SET name = $1, email = $2,birthday = $3,password = $4, updated_at = $5 WHERE user_id = $6"
	_, err := u.db.Exec(query, user.Name, user.Email, user.Birthday, user.Password, time.Now(), user.UserID)
	return err
}

func (u *UsersRepo) Create(user *modul.User) error {
	fmt.Println(user.Birthday)
	query := "INSERT INTO users (user_id,name,email,birthday,password,created_at) VALUES ($1, $2, $3,$4,$5,$6) RETURNING user_id"
	return u.db.QueryRow(query,user.UserID, user.Name, user.Email, user.Birthday, user.Password, time.Now()).Scan(&user.UserID)
}

func (u *UsersRepo) GetAll(f modul.FilterUser) ([]modul.User, error) {
	params := make(map[string]interface{})
	var args []interface{}
	filter := "WHERE true"

	if f.UserID != "" {
		params["User_id"] = f.UserID
		filter += " AND user_id = :user_id"
	}

	if f.Name != "" {
		params["name"] = f.Name
		filter += " AND name = :name"
	}
	if f.Birthday != "" {
		params["birthday"] = f.Birthday
		filter += " AND birthday = :birthday"
	}
	if f.Email != "" {
		params["email"] = f.Email
		filter += " AND email = :email"
	}

	limit := ""
	if f.Limit > 0 {
		limit = " LIMIT " + strconv.Itoa(f.Limit)
	}

	offset := ""
	if f.Offset > 0 {
		offset = " OFFSET " + strconv.Itoa(f.Offset)
	}

	query := "SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM Users " + filter + limit + offset
	query, args = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user []modul.User
	for rows.Next() {
		var Users modul.User
		if err := rows.Scan(&Users.UserID, &Users.Name, &Users.Email, &Users.Birthday, &Users.Password, &Users.CreatedAt, &Users.UpdatedAt, &Users.DeletedAt); err != nil {
			return nil, err
		}
		user = append(user, Users)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}
