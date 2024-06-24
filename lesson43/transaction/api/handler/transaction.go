package handler

import (
	"fmt"
	"transaction/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) CreateTransaction(ctx *gin.Context) {
	var transaction models.Transaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Transaction.Create(&transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) GetTransactionById(ctx *gin.Context) {
	id := ctx.Param("id")

	transaction, err := h.Transaction.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (h *handler) UpdateTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	var transaction models.Transaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Transaction.Update(id, &transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "UPDATED")
}

func (h *handler) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.Transaction.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "DELETED")
}
