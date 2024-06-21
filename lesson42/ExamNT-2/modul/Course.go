package modul

import (
	"time"
)

type Course struct {
	CourseID    string    `form:"course_id" json:"course_id"`     // Kursning xususiy identifikatori
	Title       string    `form:"title" json:"title"`             // Kurs nomi yoki sarlavhasi
	Description string    `form:"description" json:"description"` // Kurs haqida tavsif yoki ma'lumot
	CreatedAt   time.Time `form:"created_at" json:"created_at"`   // Kurs yaratilgan vaqti
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at"`   // Kurs oxirgi yangilangan vaqti
	DeletedAt   int64     `form:"deleted_at" json:"deleted_at"`   // Kurs o'chirilgan vaqti (0 bo'lsa o'chirilmagan)
}
