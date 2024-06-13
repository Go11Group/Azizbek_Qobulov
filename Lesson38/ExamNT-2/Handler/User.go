package handler

import (
	"ExamNT/modul"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(g *gin.Context) {
	user := modul.User{}

	err := g.ShouldBindBodyWithJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.User.Create(&user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Successfully created!")
}

func (h *Handler) GetUser(g *gin.Context) {
	userID := g.Param("id")
	if userID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	user, err := h.User.GetByID(userID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(g *gin.Context) {
	user := modul.User{}

	err := g.BindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.User.Update(&user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Successfully updated!")
}

func (h *Handler) DeleteUser(g *gin.Context) {
	userID := g.Param("id")
	if userID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	err := h.User.Delete(userID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, "Successfully deleted!")
}

func (h *Handler) GetAllUsers(g *gin.Context) {
	var filter modul.FilterUser
	if err := g.ShouldBindQuery(&filter); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.User.GetAll(filter)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, users)
}
