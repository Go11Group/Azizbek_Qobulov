package postgres

import (
	"ExamNT/modul"
	"database/sql"
	"strconv"
	"time"
)

type UsersRepo struct {
	db *sql.DB
}

// NewUsersRepo funksiyasi yangi UsersRepo obyektini yaratadi
func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

// GetByID funksiyasi foydalanuvchi ID bo'yicha ma'lumotni olish uchun ishlatiladi
func (u *UsersRepo) GetByID(id string) (*modul.User, string, error) {
	query := "SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1 AND deleted_at = 0"
	row := u.db.QueryRow(query, id)
	var user modul.User
	if err := row.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
		return nil, "", err
	}
	if user.DeletedAt != 0 {
		deletedTime := time.Unix(user.DeletedAt, 0).Format("02/01/2006 15:04")
		return nil, "ma'lumot " + deletedTime + " vaqtida o'chirilgan", nil
	}
	return &user, "", nil
}

// Delete funksiyasi foydalanuvchini o'chirish uchun ishlatiladi
func (u *UsersRepo) Delete(id string) error {
	query := "UPDATE users SET deleted_at = $1 WHERE user_id = $2 AND deleted_at = 0"
	_, err := u.db.Exec(query, time.Now().Unix(), id)
	return err
}

// Update funksiyasi foydalanuvchi ma'lumotlarini yangilash uchun ishlatiladi
func (u *UsersRepo) Update(user *modul.User) error {
	query := "UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, updated_at = $5 WHERE user_id = $6 AND deleted_at = 0"
	_, err := u.db.Exec(query, user.Name, user.Email, user.Birthday, user.Password, time.Now(), user.UserID)
	return err
}

// Create funksiyasi yangi foydalanuvchi yaratish uchun ishlatiladi
func (u *UsersRepo) Create(user *modul.User) error {
	query := "INSERT INTO users (user_id, name, email, birthday, password, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id"
	return u.db.QueryRow(query, user.UserID, user.Name, user.Email, user.Birthday, user.Password, time.Now()).Scan(&user.UserID)
}

// GetAll funksiyasi barcha foydalanuvchilarni filter bilan olish uchun ishlatiladi
func (u *UsersRepo) GetAll(f modul.FilterUser) ([]modul.User, error) {
	params := make(map[string]interface{})
	var args []interface{}
	filter := "WHERE deleted_at = 0"

	if f.Name != "" {
		params["name"] = f.Name
		filter += " AND name = :name"
	}

	if f.Birthday != "" {
		params["birthday"] = f.Birthday
		filter += " AND birthday = :birthday"
	}

	limit := ""
	if f.Limit > 0 {
		limit = " LIMIT " + strconv.Itoa(f.Limit)
	}

	offset := ""
	if f.Offset > 0 {
		offset = " OFFSET " + strconv.Itoa(f.Offset)
	}

	query := "SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users " + filter + limit + offset
	query, args = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []modul.User
	for rows.Next() {
		var user modul.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetByCourseID funksiyasi kurs ID bo'yicha foydalanuvchilarni olish uchun ishlatiladi
func (u *UsersRepo) GetByCourseID(courseID string) ([]modul.User, error) {
	query := `
		SELECT u.user_id, u.name, u.email, u.birthday, u.password, u.created_at, u.updated_at, u.deleted_at
		FROM users u
		INNER JOIN enrollments e ON u.user_id = e.user_id
		WHERE e.course_id = $1 AND e.deleted_at = 0`

	rows, err := u.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []modul.User
	for rows.Next() {
		var user modul.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Search funksiyasi foydalanuvchilarni izlash uchun ishlatiladi
func (u *UsersRepo) Search(filter modul.UserFilter) ([]modul.User, error) {
	params := make(map[string]interface{})
	var args []interface{}
	query := "SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE deleted_at = 0"

	if filter.Name != "" {
		params["name"] = "%" + filter.Name + "%"
		query += " AND name ILIKE :name"
	}

	if filter.Email != "" {
		params["email"] = "%" + filter.Email + "%"
		query += " AND email ILIKE :email"
	}

	if filter.AgeFrom > 0 {
		yearTo := time.Now().AddDate(-filter.AgeFrom, 0, 0)
		params["yearTo"] = yearTo
		query += " AND birthday <= :yearTo"
	}

	if filter.AgeTo > 0 {
		yearFrom := time.Now().AddDate(-(filter.AgeTo + 1), 0, 0)
		params["yearFrom"] = yearFrom
		query += " AND birthday >= :yearFrom"
	}
	query, args = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []modul.User
	for rows.Next() {
		var user modul.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
