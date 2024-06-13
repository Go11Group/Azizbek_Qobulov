package postgres

import (
	"ExamNT/modul"
	"database/sql"
	"strconv"
	"time"
)

type CourseRepo struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

func (u *CourseRepo) GetByID(id string) (*modul.Course, error) {
	query := "SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses WHERE course_id = $1"
	row := u.db.QueryRow(query,id)
	var course modul.Course
	if err := row.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt); err != nil {
		return nil, err
	}
	return &course, nil
}

func (u *CourseRepo) Delete(id string) error {
	query := "UPDATE courses SET deleted_at = $1 WHERE course_id = $2"
	_, err := u.db.Exec(query, time.Now().Unix(), id)
	return err
}

func (u *CourseRepo) Update(course *modul.Course) error {
	query := "UPDATE courses SET title = $1, description = $2, updated_at =$3 WHERE course_id = $4"
	_, err := u.db.Exec(query, course.Title, course.Description, time.Now(), course.CourseID)
	return err
}

func (u *CourseRepo) Create(course *modul.Course) error {
	query := "INSERT INTO courses (course_id,title, description,created_at) VALUES ($1, $2, $3, $4) RETURNING course_id"
	return u.db.QueryRow(query,course.CourseID, course.Title, course.Description, time.Now()).Scan(&course.CourseID)
}

func (u *CourseRepo) GetAll(f modul.FilterCourse) ([]modul.Course, error) {
	params := make(map[string]interface{})
	var args []interface{}
	filter := "WHERE true"

	if f.Title != "" {
		params["title"] = f.Title
		filter += " AND title = :title"
	}

	if f.CourseID != "" {
		params["course_id"] = f.CourseID
		filter += " AND course_id = :course_id"
	}

	limit := ""
	if f.Limit > 0 {
		limit = " LIMIT " + strconv.Itoa(f.Limit)
	}

	offset := ""
	if f.Offset > 0 {
		offset = " OFFSET " + strconv.Itoa(f.Offset)
	}

	query := "SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses " + filter + limit + offset
	query, args = ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []modul.Course
	for rows.Next() {
		var course modul.Course
		if err := rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
