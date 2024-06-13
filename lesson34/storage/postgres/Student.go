<<<<<<< HEAD
package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{db: db}
}

func (r *StudentRepo) GetById(id string) (*model.Student, error) {
	query := "SELECT id, name, phone, email FROM students WHERE id = $1"
	row := r.db.QueryRow(query, id)
	var student model.Student
	if err := row.Scan(&student.ID, &student.Name, &student.Phone, &student.Email); err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepo) Delete(id string) error {
	query := "DELETE FROM students WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *StudentRepo) Update(student *model.Student) error {
	query := "UPDATE students SET name = $1, phone = $2, email = $3 WHERE id = $4"
	_, err := r.db.Exec(query, student.Name, student.Phone, student.Email, student.ID)
	return err
}

func (r *StudentRepo) Create(student *model.Student) error {
	query := "INSERT INTO students (name, phone, email) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, student.Name, student.Phone, student.Email).Scan(&student.ID)
}
=======
package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{db: db}
}

func (b *StudentRepo) GetById(id string) (model.Student, error) {
	var student model.Student
	err := b.db.QueryRow("SELECT * FROM student WHERE Student_ID = $1", id).Scan(
		&student.ID,
		&student.Name,
		&student.Email,
		&student.Birth_year,
		&student.Phone,
	)

	return student, err
}

func (b *StudentRepo) Delete(id string) error {

	_, err := b.db.Exec(`DELETE FROM student WHERE student_id = $1`, id)
	return err
}

func (b *StudentRepo) Update(student *model.Student) error {

	_, err := b.db.Exec(`
	UPDATE student
	SET Name = $1,
		email = $2,
		birth_year = $3,
		phone = $4
	WHERE id = $6`,
		student.Name,
		student.Email,
		student.Birth_year,
		student.Phone)

	return err
}
>>>>>>> origin/main
