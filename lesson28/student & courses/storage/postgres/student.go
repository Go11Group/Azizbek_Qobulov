package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson28/model"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.Student, error) {
	rows, err := u.Db.Query(`select s.id, s.name, age, gender, c.name from students s
					left join courses c on c.id = s.course_id `)
	if err != nil {
		return nil, err
	}

	var users []model.Student
	var user model.Student
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByID(id string) (model.Student, error) {
	var user model.Student

	err := u.Db.QueryRow(`select s.id, s.name, age, gender, c.name from students s
					left join course c on c.id = s.course_id where s.id = $1`, id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	if err != nil {
		return model.Student{}, err
	}

	return user, nil
}

func Create(db *sql.DB, user model.Student) error {

	//uuid.NewString()

	_,err := db.Exec(`INSERT INTO Student(id,Name,Age,Gender,Course) values($1,$2,$3,$4,$5)`,user.ID, user.Name, user.Age, user.Gender, user.Course)
	if err != nil {
		return err
	}
	return nil
}

func Update(db *sql.DB, user model.Student) error {

	_,err := db.Exec(`UPDATE Student set Name = $1, Age = $2, Gender = $3, Course = $4`,user.ID, user.Name, user.Age, user.Gender, user.Course)
	if err != nil {
		return err
	}
	return nil
}

func Delete(db *sql.DB, id string) error {
	_,err := db.Exec(`DELETE FROM Student where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
