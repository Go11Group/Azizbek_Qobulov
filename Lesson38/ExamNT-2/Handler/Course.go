package handler

import (
	"ExamNT/modul"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCourse funksiyasi yangi kurs yaratadi
func (h *Handler) CreateCourse(g *gin.Context) {
	course := modul.Course{}

	// JSON formatdagi ma'lumotlarni structga bog'laydi
	err := g.BindJSON(&course)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Yangi kursni bazaga qo'shadi
	err = h.Course.Create(&course)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kurs yaratildi deb javob qaytaradi
	g.JSON(http.StatusOK, course)
}

// GetCourse funksiyasi kurs ID bo'yicha kursni qaytaradi
func (h *Handler) GetCourse(g *gin.Context) {
	courseID, hasID := g.Params.Get("id")
	if !hasID {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing course ID"})
		return
	}

	// Kursni ID bo'yicha olish
	course, err := h.Course.GetByID(courseID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kurs ma'lumotlarini qaytarish
	g.JSON(http.StatusOK, course)
}

// UpdateCourse funksiyasi kursni yangilaydi
func (h *Handler) UpdateCourse(g *gin.Context) {
	course := modul.Course{}

	// JSON formatdagi ma'lumotlarni structga bog'laydi
	err := g.BindJSON(&course)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kursni yangilash
	err = h.Course.Update(&course)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Yangilangan kursni qaytarish
	g.JSON(http.StatusOK, course)
}

// DeleteCourse funksiyasi kursni o'chiradi
func (h *Handler) DeleteCourse(g *gin.Context) {
	courseID := g.Param("id")
	if courseID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing course ID"})
		return
	}

	// Kursni ID bo'yicha o'chirish
	err := h.Course.Delete(courseID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// O'chirildi deb javob qaytarish
	g.JSON(http.StatusOK, "Successfully deleted!")
}

// GetAllCourses funksiyasi barcha kurslarni filter bilan qaytaradi
func (h *Handler) GetAllCourses(g *gin.Context) {
	var filter modul.FilterCourse
	// Query parametrlari orqali filterni olish
	if err := g.ShouldBindQuery(&filter); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Filter bo'yicha kurslarni olish
	courses, err := h.Course.GetAll(filter)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kurslarni qaytarish
	g.JSON(http.StatusOK, courses)
}

// GetCoursesByUser funksiyasi foydalanuvchi ID bo'yicha kurslarni qaytaradi
func (h *Handler) GetCoursesByUser(c *gin.Context) {
	userID, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Foydalanuvchi ID bo'yicha kurslarni olish
	courses, err := h.Course.GetCourseByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kurslarni qaytarish
	c.JSON(http.StatusOK, courses)
}

// GetPopularCourses funksiyasi eng ommabop kurslarni belgilangan muddat davomida qaytaradi
func (h *Handler) GetPopularCourses(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	// Belgilangan muddatdagi kurslarni olish
	popularCourses, err := h.Course.GetPopularCourse(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Eng ommabop kurslarni va belgilangan muddatni qaytarish
	c.JSON(http.StatusOK, gin.H{
		"time_period": gin.H{
			"start_date": startDate,
			"end_date":   endDate,
		},
		"popular_courses": popularCourses,
	})
}
