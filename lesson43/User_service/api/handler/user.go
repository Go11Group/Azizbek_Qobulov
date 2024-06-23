package handler

import (
	"fmt"
	"user_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateUser(ctx *gin.Context) {
	var usr models.CreateUser
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.User.Create(&usr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.User.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}
