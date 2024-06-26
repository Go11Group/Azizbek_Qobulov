package handler

import (
	"fmt"
	"user_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateUser(ctx *gin.Context) {
	var usr models.User
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

func (h *handler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var usr models.User
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.User.Update(id, &usr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "UPDATED")
}

func (h *handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.User.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "DELETED")
}
