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
