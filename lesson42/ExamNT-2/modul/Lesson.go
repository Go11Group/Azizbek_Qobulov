package modul

import (
	"time"
)

type Lesson struct {
	LessonID  string    `form:"lesson_id" json:"lesson_id"`   // Dars identifikatori
	CourseID  string    `form:"course_id" json:"course_id"`   // Kurs identifikatori
	Title     string    `form:"title" json:"title"`           // Dars sarlavhasi
	Content   string    `form:"content" json:"content"`       // Dars mazmuni
	CreatedAt time.Time `form:"created_at" json:"created_at"` // Dars yaratilgan vaqti
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"` // Dars yangilangan vaqti
	DeletedAt int64     `form:"deleted_at" json:"deleted_at"` // Dars o'chirilgan vaqti (0 o'chirilmagan)
}
