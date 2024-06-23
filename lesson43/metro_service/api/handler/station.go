package handler

import (
	"fmt"
	"net/http"

	"github.com/Go11Group/at_lesson/lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateStation(ctx *gin.Context) {
	var stn models.Station
	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Station.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) UpdateCard(ctx *gin.Context) {
	id := ctx.Param("id")
	var card models.Station
	err := ctx.ShouldBindJSON(&card)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Station.Update(id, &card)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "UPDATED")
}

func (h *handler) DeleteCard(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.Station.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "DELETED")
}