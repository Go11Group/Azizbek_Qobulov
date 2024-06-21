package handler

import (
	"ExamNT/modul"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser funksiyasi yangi foydalanuvchi yaratadi
func (h *Handler) CreateUser(g *gin.Context) {
	var user modul.User

	// JSON formatdagi ma'lumotlarni structga bog'laydi
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Foydalanuvchi ma'lumotlarini bazaga qo'shadi
	if err := h.User.Create(&user); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Yangi yaratilgan foydalanuvchini qaytaradi
	g.JSON(http.StatusOK, user)
}

// GetUser funksiyasi foydalanuvchini ID bo'yicha qaytaradi
func (h *Handler) GetUser(g *gin.Context) {
	userID := g.Param("id")
	if userID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	// Foydalanuvchi ma'lumotlarini ID bo'yicha olish
	user, message, err := h.User.GetByID(userID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if message != "" {
		g.JSON(http.StatusOK, gin.H{"message": message})
		return
	}

	// Foydalanuvchi ma'lumotlarini qaytarish
	g.JSON(http.StatusOK, user)
}

// UpdateUser funksiyasi foydalanuvchi ma'lumotlarini yangilaydi
func (h *Handler) UpdateUser(g *gin.Context) {
	var user modul.User

	// JSON formatdagi ma'lumotlarni structga bog'laydi
	if err := g.BindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Foydalanuvchi ma'lumotlarini yangilash
	if err := h.User.Update(&user); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Yangilangan foydalanuvchi ma'lumotlarini qaytarish
	g.JSON(http.StatusOK, user)
}

// DeleteUser funksiyasi foydalanuvchi ma'lumotlarini o'chiradi
func (h *Handler) DeleteUser(g *gin.Context) {
	userID, hasID := g.Params.Get("id")
	if !hasID {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	// Foydalanuvchi ma'lumotlarini ID bo'yicha o'chirish
	if err := h.User.Delete(userID); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// O'chirildi deb javob qaytarish
	g.JSON(http.StatusOK, gin.H{"message": "Successfully deleted!"})
}

// GetAllUsers funksiyasi barcha foydalanuvchilarni filter bilan qaytaradi
func (h *Handler) GetAllUsers(g *gin.Context) {
	var filter modul.FilterUser
	// Query parametrlari orqali filterni olish
	if err := g.ShouldBindQuery(&filter); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Filter bo'yicha barcha foydalanuvchilarni olish
	users, err := h.User.GetAll(filter)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Foydalanuvchi ma'lumotlarini qaytarish
	g.JSON(http.StatusOK, users)
}

// GetEnrolledUsersByCourse funksiyasi kursga yozilgan foydalanuvchilarni qaytaradi
func (h *Handler) GetEnrolledUsersByCourse(c *gin.Context) {
	courseID, hasID := c.Params.Get("id")
	if !hasID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Kurs ID bo'yicha yozilgan foydalanuvchilarni olish
	users, err := h.User.GetByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Foydalanuvchi ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, users)
}

// SearchUsers funksiyasi foydalanuvchilarni qidirish uchun ishlatiladi
func (h *Handler) SearchUsers(c *gin.Context) {
	var filter modul.UserFilter
	// Query parametrlari orqali filterni olish
	filter.Name = c.Query("name")
	filter.Email = c.Query("email")
	filter.AgeFrom, _ = strconv.Atoi(c.Query("age_from"))
	filter.AgeTo, _ = strconv.Atoi(c.Query("age_to"))

	// Filter bo'yicha foydalanuvchilarni qidirish
	users, err := h.User.Search(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Foydalanuvchi ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, users)
}
