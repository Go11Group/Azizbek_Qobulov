package postgres

import (
	"ExamNT/modul"
	"database/sql"
	"strconv"
	"time"
)

type EnrollmentRepo struct {
	db *sql.DB
}

// NewEnrollmentRepo yangi EnrollmentRepo obyektini yaratadi
func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db: db}
}

// Enroll funksiyasi yangi enrollments yaratadi
func (r *EnrollmentRepo) Enroll(enrollment *modul.Enrollment) error {
	query := `INSERT INTO enrollments (enrollment_id, user_id, course_id, enrollment_date, created_at)
              VALUES ($1, $2, $3, $4, $5) RETURNING enrollment_id`
	return r.db.QueryRow(query, enrollment.EnrollmentID, enrollment.UserID, enrollment.CourseID, time.Now(), time.Now()).Scan(&enrollment.EnrollmentID)
}

// GetByID funksiyasi enrollmentni ID bo'yicha qaytaradi
func (r *EnrollmentRepo) GetByID(id string) (*modul.Enrollment, string, error) {
	query := "SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at FROM enrollments WHERE enrollment_id = $1"
	row := r.db.QueryRow(query, id)
	var enrollment modul.Enrollment
	if err := row.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt); err != nil {
		return nil, "", err
	}
	if enrollment.DeletedAt != 0 {
		deletedTime := time.Unix(enrollment.DeletedAt, 0).Format("02/01/2006 15:04")
		return nil, "data deleted at " + deletedTime, nil
	}
	return &enrollment, "", nil
}

// Update funksiyasi enrollmentni yangilaydi
func (r *EnrollmentRepo) Update(enrollment *modul.Enrollment) error {
	query := `UPDATE enrollments SET user_id = $1, course_id = $2, enrollment_date = $3, updated_at = $4 WHERE enrollment_id = $5`
	_, err := r.db.Exec(query, enrollment.UserID, enrollment.CourseID, enrollment.EnrollmentDate, time.Now(), enrollment.EnrollmentID)
	return err
}

// Delete funksiyasi enrollmentni ID bo'yicha o'chiradi
func (r *EnrollmentRepo) Delete(id string) error {
	query := `UPDATE enrollments SET deleted_at = $1 WHERE enrollment_id = $2 and deleted_at = 0`
	_, err := r.db.Exec(query, time.Now().Unix(), id)
	return err
}

// GetAll funksiyasi barcha enrollmentsni filter bilan qaytaradi
func (r *EnrollmentRepo) GetAll(f *modul.FilterEnrollment) ([]modul.Enrollment, error) {
	var enrollments []modul.Enrollment
	query := `SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at
              FROM enrollments WHERE deleted_at = 0`
	params := make(map[string]interface{})
	var args []interface{}

	if f.UserID != "" {
		params["user_id"] = f.UserID
		query += " AND user_id = :user_id"
	}

	if f.CourseID != "" {
		params["course_id"] = f.CourseID
		query += " AND course_id = :course_id"
	}

	limit := ""
	if f.Limit > 0 {
		limit = " LIMIT " + strconv.Itoa(f.Limit)
	}

	offset := ""
	if f.Offset > 0 {
		offset = " OFFSET " + strconv.Itoa(f.Offset)
	}

	query += limit + offset
	query, args = ReplaceQueryParams(query, params)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var enrollment modul.Enrollment
		if err := rows.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt); err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}
