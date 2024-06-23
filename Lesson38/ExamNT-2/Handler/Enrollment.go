package handler

import (
	"ExamNT/modul"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EnrollUser funksiyasi foydalanuvchini kursga ro'yxatdan o'tkazadi
func (h *Handler) EnrollUser(c *gin.Context) {
	var enrollment modul.Enrollment
	// JSON formatdagi ma'lumotlarni structga bog'laydi
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ro'yxatdan o'tish ma'lumotlarini bazaga qo'shadi
	if err := h.Enrollment.Enroll(&enrollment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ro'yxatdan o'tgan foydalanuvchini qaytaradi
	c.JSON(http.StatusOK, enrollment)
}

// GetEnrollmentByID funksiyasi ro'yxatdan o'tish ID bo'yicha ma'lumotlarni qaytaradi
func (h *Handler) GetEnrollmentByID(g *gin.Context) {
	enrollmentID := g.Param("id")
	if enrollmentID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing enrollment ID"})
		return
	}

	// Ro'yxatdan o'tish ma'lumotlarini ID bo'yicha olish
	enrollment, message, err := h.Enrollment.GetByID(enrollmentID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if message != "" {
		g.JSON(http.StatusOK, gin.H{"message": message})
		return
	}

	// Ro'yxatdan o'tish ma'lumotlarini qaytarish
	g.JSON(http.StatusOK, enrollment)
}

// UpdateEnrollment funksiyasi ro'yxatdan o'tish ma'lumotlarini yangilaydi
func (h *Handler) UpdateEnrollment(c *gin.Context) {
	var enrollment modul.Enrollment
	// JSON formatdagi ma'lumotlarni structga bog'laydi
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ro'yxatdan o'tish ma'lumotlarini yangilash
	if err := h.Enrollment.Update(&enrollment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Yangilangan ro'yxatdan o'tish ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, enrollment)
}

// DeleteEnrollment funksiyasi ro'yxatdan o'tish ma'lumotlarini o'chiradi
func (h *Handler) DeleteEnrollment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Ro'yxatdan o'tish ma'lumotlarini ID bo'yicha o'chirish
	if err := h.Enrollment.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// O'chirildi deb javob qaytarish
	c.JSON(http.StatusOK, "Successfully deleted!")
}

// GetAllEnrollments funksiyasi barcha ro'yxatdan o'tish ma'lumotlarini filter bilan qaytaradi
func (h *Handler) GetAllEnrollments(c *gin.Context) {
	filter := modul.FilterEnrollment{}
	// Query parametrlari orqali filterni olish
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Filter bo'yicha barcha ro'yxatdan o'tish ma'lumotlarini olish
	enrollments, err := h.Enrollment.GetAll(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ro'yxatdan o'tish ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, enrollments)
}
