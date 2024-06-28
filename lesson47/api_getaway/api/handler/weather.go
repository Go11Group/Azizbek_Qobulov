package handler

import (
	pw "api/genproto/WeatherService"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ReportWeather(ctx *gin.Context) {
	r := pw.WeatherConditionRequest{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "error while decoding request body",
		})
		log.Println("error while decoding request body ", err)
		return
	}
	res, err := h.Weather.ReportWeatherCondition(context.Background(), &r)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "StatusInternalServerError",
			"message": err.Error(),
		})
		log.Println("error while reporting weather condition", err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetCurrentWeather(ctx *gin.Context) {
	city := ctx.Query("city")
	if len(city) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "City not found",
		})
		log.Println("City not found")
		return
	}
	req := pw.CurrentWeatherRequest{Location: city}
	curWeather, err := h.Weather.GetCurrentWeather(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "StatusInternalServerError",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, curWeather)
}

func (h *Handler) GetWeatherForecast(ctx *gin.Context) {
	city := ctx.Query("city")
	if len(city) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "StatusBadRequest",
			"message": "City not found",
		})
		log.Println("City not found")
		return
	}
	req := pw.WeatherForecastRequest{Location: city}
	forecast, err := h.Weather.GetWeatherForecast(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "StatusInternalServerError",
			"message": err.Error(),
		})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, forecast)
}
