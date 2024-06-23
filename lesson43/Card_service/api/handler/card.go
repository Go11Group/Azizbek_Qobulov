package handler

import (
	"fmt"
	"Card_service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateCard(ctx *gin.Context) {
	var card models.Card
	err := ctx.ShouldBindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Card.Create(&card)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) GetCardById(ctx *gin.Context) {
	id := ctx.Param("id")

	card, err := h.Card.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, card)
}

func (h *handler) UpdateCard(ctx *gin.Context) {
	id := ctx.Param("id")
	var card models.Card
	err := ctx.ShouldBindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Card.Update(id, &card)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "UPDATED")
}

func (h *handler) DeleteCard(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.Card.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "DELETED")
}
