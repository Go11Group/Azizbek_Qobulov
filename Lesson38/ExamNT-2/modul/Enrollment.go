package modul

import (
	"time"
)

type Enrollment struct {
	EnrollmentID   string    `form:"enrollment_id" json:"enrollment_id"`     // Yozilish identifikatori
	UserID         string    `form:"user_id" json:"user_id"`                 // Foydalanuvchi identifikatori
	CourseID       string    `form:"course_id" json:"course_id"`             // Kurs identifikatori
	EnrollmentDate time.Time `form:"enrollment_date" json:"enrollment_date"` // Kursga yozilish sanasi
	CreatedAt      time.Time `form:"created_at" json:"created_at"`           // Yozilgan vaqti
	UpdatedAt      time.Time `form:"updated_at" json:"updated_at"`           // Oxirgi yangilangan vaqti
	DeletedAt      int64     `form:"deleted_at" json:"deleted_at"`           // O'chirilgan vaqti (0 bo'lsa o'chirilmagan)
}
