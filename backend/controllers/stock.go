package controllers

import (
	"github.com/Raajheer1/hackutd-ix/m/v2/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StockInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	Ticker string `json:"ticker" binding:"required"`
	Shares string `json:"shares" binding:"required"`
}

func CreateStock(c *gin.Context) {
	var input StockInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Stock{}
	s.UserID = input.UserID
	s.Ticker = input.Ticker
	s.Shares = input.Shares

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": authToken})
}
