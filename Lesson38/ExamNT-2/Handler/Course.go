package handler

import (
	"ExamNT/modul"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourse(g *gin.Context) {
	course := modul.Course{}

	err := g.BindJSON(&course)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Course.Create(&course)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Successfully created!")
}

func (h *Handler) GetCourse(g *gin.Context) {
	courseID := g.Param("id")
	if courseID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing course ID"})
		return
	}

	course, err := h.Course.GetByID(courseID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, course)
}

func (h *Handler) UpdateCourse(g *gin.Context) {
	course := modul.Course{}

	err := g.BindJSON(&course)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.Course.Update(&course)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Successfully updated!")
}

func (h *Handler) DeleteCourse(g *gin.Context) {
	courseID := g.Param("id")
	if courseID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing course ID"})
		return
	}

	err := h.Course.Delete(courseID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Successfully deleted!")
}

func (h *Handler) GetAllCourses(g *gin.Context) {
	var filter modul.FilterCourse
	if err := g.ShouldBindQuery(&filter); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courses, err := h.Course.GetAll(filter)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, courses)
}
