package modul

import (
	"time"
)

type User struct {
	UserID    string    `form:"user_id" json:"user_id"`       // Foydalanuvchi identifikatori
	Name      string    `form:"name" json:"name"`             // Foydalanuvchi ismi
	Email     string    `form:"email" json:"email"`           // Foydalanuvchi elektron pochta manzili
	Birthday  string    `form:"birthday" json:"birthday"`     // Foydalanuvchi tug'ilgan kun sanasi
	Password  string    `form:"password" json:"password"`     // Foydalanuvchi paroli
	CreatedAt time.Time `form:"created_at" json:"created_at"` // Foydalanuvchi yaratilgan vaqti
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"` // Foydalanuvchi ma'lumotlari yangilangan vaqti
	DeletedAt int64     `form:"deleted_at" json:"deleted_at"` // Foydalanuvchi o'chirilgan vaqti (0 o'chirilmagan)
}
