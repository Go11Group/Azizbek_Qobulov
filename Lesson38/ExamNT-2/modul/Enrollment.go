package modul

import (
	"time"
)

type Enrollment struct {
	EnrollmentID   string    `json:"enrollment_id"`
	UserID         string    `json:"user_id"`
	CourseID       string    `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      int64     `json:"deleted_at"`
}
