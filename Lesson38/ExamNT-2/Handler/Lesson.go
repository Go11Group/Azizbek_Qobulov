package handler

import (
	"ExamNT/modul"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLesson(c *gin.Context) {
	var lesson modul.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Lesson.Create(&lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Successfully created!")
}

func (h *Handler) GetLessonByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {  // Corrected condition to check for empty ID
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	lesson, err := h.Lesson.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) GetAllLessons(c *gin.Context) {
	lesson := modul.FilterLesson{}

	if err := c.ShouldBindQuery(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lessons, err := h.Lesson.GetAll(&lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lessons)
}

func (h *Handler) UpdateLesson(c *gin.Context) {
	var lesson modul.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Lesson.Update(&lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Successfully updated!")
}

func (h *Handler) DeleteLesson(c *gin.Context) {
	id := c.Param("id")
	if id == "" {  // Corrected condition to check for empty ID
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.Lesson.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted!")
}
