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

// NewCourseRepo yangi CourseRepo obyektini yaratadi
func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

// GetByID funksiyasi kursni ID bo'yicha qaytaradi
func (u *CourseRepo) GetByID(id string) (*modul.Course, error) {
	query := "SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses WHERE course_id = $1 and deleted_at = 0"
	row := u.db.QueryRow(query, id)
	var course modul.Course
	if err := row.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt); err != nil {
		return nil, err
	}
	return &course, nil
}

// Delete funksiyasi kursni ID bo'yicha o'chiradi
func (u *CourseRepo) Delete(id string) error {
	query := "UPDATE courses SET deleted_at = $1 WHERE course_id = $2 and deleted_at = 0"
	_, err := u.db.Exec(query, time.Now().Unix(), id)
	return err
}

// Update funksiyasi kursni yangilaydi
func (u *CourseRepo) Update(course *modul.Course) error {
	query := "UPDATE courses SET title = $1, description = $2, updated_at =$3 WHERE course_id = $4 and deleted_at = 0"
	_, err := u.db.Exec(query, course.Title, course.Description, time.Now(), course.CourseID)
	return err
}

// Create funksiyasi yangi kurs yaratadi
func (u *CourseRepo) Create(course *modul.Course) error {
	query := "INSERT INTO courses (course_id, title, description, created_at) VALUES ($1, $2, $3, $4) RETURNING course_id"
	return u.db.QueryRow(query, course.CourseID, course.Title, course.Description, time.Now()).Scan(&course.CourseID)
}

// GetAll funksiyasi barcha kurslarni filter bilan qaytaradi
func (u *CourseRepo) GetAll(f modul.FilterCourse) ([]modul.Course, error) {
	params := make(map[string]interface{})
	var args []interface{}
	filter := "WHERE deleted_at = 0"

	if f.Title != "" {
		params["title"] = f.Title
		filter += " AND title = :title"
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

// GetCourseByUserID funksiyasi foydalanuvchi ID bo'yicha kurslarni qaytaradi
func (u *CourseRepo) GetCourseByUserID(userID string) ([]modul.Course, error) {
	query := `
		SELECT c.course_id, c.title, c.description, c.created_at, c.updated_at, c.deleted_at
		FROM courses c
		INNER JOIN enrollments e ON c.course_id = e.course_id
		WHERE e.user_id = $1 AND e.deleted_at = 0`

	rows, err := u.db.Query(query, userID)
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

	return courses, nil
}

// GetPopularCourse funksiyasi berilgan sana oralig'ida eng mashhur kurslarni qaytaradi
func (r *CourseRepo) GetPopularCourse(startDate, endDate string) ([]modul.PopularCourse, error) {
	query := `SELECT 
			c.course_id, 
			c.title AS course_title,
			COUNT(e.enrollment_id) AS enrollments_count
			FROM courses c
			INNER JOIN enrollments e ON c.course_id = e.course_id
			WHERE e.enrollment_date BETWEEN $1::date AND $2::date AND c.deleted_at = 0
			GROUP BY c.course_id, c.title
			ORDER BY enrollments_count DESC`

	rows, err := r.db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var popularCourses []modul.PopularCourse
	for rows.Next() {
		var pc modul.PopularCourse
		if err := rows.Scan(&pc.CourseID, &pc.CourseTitle, &pc.EnrollmentsCount); err != nil {
			return nil, err
		}
		popularCourses = append(popularCourses, pc)
	}

	return popularCourses, nil
}
