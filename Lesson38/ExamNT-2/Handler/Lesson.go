package handler

import (
	"ExamNT/modul"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateLesson funksiyasi yangi dars yaratadi
func (h *Handler) CreateLesson(c *gin.Context) {
	var lesson modul.Lesson
	// JSON formatdagi ma'lumotlarni structga bog'laydi
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dars ma'lumotlarini bazaga qo'shadi
	if err := h.Lesson.Create(&lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Yangi yaratilgan darsni qaytaradi
	c.JSON(http.StatusOK, lesson)
}

// GetLessonByID funksiyasi darsni ID bo'yicha qaytaradi
func (h *Handler) GetLessonByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Dars ma'lumotlarini ID bo'yicha olish
	lesson, message, err := h.Lesson.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if message != "" {
		c.JSON(http.StatusOK, gin.H{"message": message})
		return
	}

	if lesson == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	// Dars ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, lesson)
}

// GetAllLessons funksiyasi barcha darslarni filter bilan qaytaradi
func (h *Handler) GetAllLessons(c *gin.Context) {
	var filter modul.FilterLesson

	// Query parametrlari orqali filterni olish
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Filter bo'yicha barcha darslarni olish
	lessons, err := h.Lesson.GetAll(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Dars ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, lessons)
}

// UpdateLesson funksiyasi dars ma'lumotlarini yangilaydi
func (h *Handler) UpdateLesson(c *gin.Context) {
	var lesson modul.Lesson
	// JSON formatdagi ma'lumotlarni structga bog'laydi
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dars ma'lumotlarini yangilash
	if err := h.Lesson.Update(&lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Yangilangan dars ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, lesson)
}

// DeleteLesson funksiyasi dars ma'lumotlarini o'chiradi
func (h *Handler) DeleteLesson(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Dars ma'lumotlarini ID bo'yicha o'chirish
	if err := h.Lesson.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// O'chirildi deb javob qaytarish
	c.JSON(http.StatusOK, "Successfully deleted!")
}

// GetLessonsByCourse funksiyasi kursga tegishli barcha darslarni qaytaradi
func (h *Handler) GetLessonsByCourse(c *gin.Context) {
	courseID, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Kurs ID bo'yicha barcha darslarni olish
	lessons, err := h.Lesson.GetByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Dars ma'lumotlarini qaytarish
	c.JSON(http.StatusOK, lessons)
}
