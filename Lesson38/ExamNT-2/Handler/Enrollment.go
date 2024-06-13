package handler

import (
    "ExamNT/modul"
    "net/http"

    "github.com/gin-gonic/gin"
)

func (h *Handler) EnrollUser(c *gin.Context) {
    var enrollment modul.Enrollment
    if err := c.ShouldBindJSON(&enrollment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.Enrollment.Enroll(&enrollment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, "Successfully enrolled!")
}

func (h *Handler) GetEnrollmentByID(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    enrollment, err := h.Enrollment.GetByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, enrollment)
}

func (h *Handler) DeleteEnrollment(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := h.Enrollment.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, "Successfully deleted!")
}

func (h *Handler) GetAllEnrollments(c *gin.Context) {
    filter := modul.FilterEnrollment{}
    if err := c.ShouldBindQuery(&filter); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    enrollments, err := h.Enrollment.GetAll(&filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, enrollments)
}
