package main

import (
	"errors"
	"net/http"
	"rate-api/service"

	"github.com/gin-gonic/gin"
)

func getRate(context *gin.Context) {
	rate, err := service.GetRateFromBinance()
	if err != nil {
		context.Writer.WriteHeader(500)
		return
	}
	context.IndentedJSON(http.StatusOK, rate.Price)

}

func subscribe(context *gin.Context) {
	var newEmail service.Email

	if err := context.Bind(&newEmail); err != nil {
		http.Error(context.Writer, err.Error(), 400)
		return
	}

	if err := service.AddEmail(newEmail); err != nil {
		if errors.Is(err, service.ErrEmailSubscribed) {
			http.Error(context.Writer, err.Error(), 409)
			return
		}
		if errors.Is(err, service.ErrEmailNotValid) {
			http.Error(context.Writer, err.Error(), 400)
			return
		}

		http.Error(context.Writer, err.Error(), 500)
		return
	}

	context.Status(http.StatusCreated)

}

func sendEmails(context *gin.Context) {
	if err := service.SendEmails(); err != nil {
		http.Error(context.Writer, "Failed to send emails", 500)
		return
	}
	context.Status(http.StatusOK)
}

func main() {
	router := gin.Default()
	router.GET("/rate", getRate)
	router.POST("/subscribe", subscribe)
	router.POST("/sendEmails", sendEmails)
	router.Run("0.0.0.0:8080")
}
