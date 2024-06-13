package postgres

import (
    "ExamNT/modul"
    "database/sql"
    "time"
    "strconv"
)

type EnrollmentRepo struct {
    db *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
    return &EnrollmentRepo{db: db}
}

func (r *EnrollmentRepo) Enroll(enrollment *modul.Enrollment) error {
    query := `INSERT INTO enrollments (user_id, course_id, enrollment_date, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5) RETURNING enrollment_id`
    return r.db.QueryRow(query, enrollment.UserID, enrollment.CourseID, time.Now(), time.Now(), time.Now()).Scan(&enrollment.EnrollmentID)
}

func (r *EnrollmentRepo) GetByID(id string) (*modul.Enrollment, error) {
    enrollment := &modul.Enrollment{}
    query := `SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at
              FROM enrollments WHERE enrollment_id = $1`
    err := r.db.QueryRow(query, id).Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
    return enrollment, err
}

func (r *EnrollmentRepo) Delete(id string) error {
    query := `UPDATE enrollments SET deleted_at = $1 WHERE enrollment_id = $2`
    _, err := r.db.Exec(query, time.Now(), id)
    return err
}

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
