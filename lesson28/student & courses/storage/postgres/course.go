package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson28/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (u *CourseRepo) GetAllCourses() ([]model.Course, error) {
	rows, err := u.DB.Query(`select id, name,field from Courses`)
	if err != nil {
		return nil, err
	}
	var kurslar []model.Course
	var kusr model.Course
	for rows.Next() {
		err = rows.Scan(&kusr.ID, &kusr.Name, &kusr.Field)
		if err != nil {
			return nil, err
		}
		kurslar = append(kurslar, kusr)
	}

	return kurslar, nil
}

func (u *CourseRepo) GetByID(id string) (model.Course, error) {
	var kurs model.Course

	err := u.DB.QueryRow(`select id Name, Field from Courses where id = $1`, id).
		Scan(&kurs.ID, &kurs.Name, &kurs.Field)
	if err != nil {
		return model.Course{}, err
	}

	return kurs, nil
}

func (c *CourseRepo) Create(db *sql.DB, course model.Course) error {
	_,err := db.Exec(`INSERT INTO Course(id,Name,Field) values($1,$2,$3)`,course.ID, course.Name, course.Field)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo)Update(db *sql.DB, course model.Course) error {

	_,err := db.Exec(`UPDATE Course set Name = $1, Field = $2`,course.ID, course.Name,course.Field)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo)Delete(db *sql.DB, id string) error {
	_,err := db.Exec(`DELETE FROM Course where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}