package handler

import (
	"net/http"
	"strconv"

	"github.com/Go11Group/at_lesson/lesson34/model"
	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Talaba *postgres.StudentRepo
}

func NewHandler(repo *postgres.StudentRepo) *Handler {
	return &Handler{Talaba: repo}
}

func (h *Handler) SetupRoutes(r *gin.Engine) {
	r.GET("/student/:id", h.StudentGetById)
	r.DELETE("/student/:id", h.StudentDeleteById)
	r.PUT("/student/:id", h.StudentUpdateById)
	r.POST("/student", h.StudentCreate)
}

func (h *Handler) StudentGetById(c *gin.Context) {
	id := c.Param("id")

	student, err := h.Talaba.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *Handler) StudentDeleteById(c *gin.Context) {
	id := c.Param("id")

	err := h.Talaba.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Success to delete student")
}

func (h *Handler) StudentUpdateById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.ID = id

	err = h.Talaba.Update(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Success to update")
}

func (h *Handler) StudentCreate(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Talaba.Create(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}
