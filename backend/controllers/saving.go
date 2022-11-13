package controllers

import (
	"github.com/Raajheer1/hackutd-ix/m/v2/models"
	"github.com/Raajheer1/hackutd-ix/m/v2/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SavingInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	Amount string `json:"amount" binding:"required"`
}

func CreateSaving(c *gin.Context) {
	var input SavingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := models.Saving{}
	v.UserID = input.UserID
	v.Amount = input.Amount

	_, err := v.CreateSaving()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Saving Added"})
}

func GetSaving(c *gin.Context) {
	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savings, err := models.GetSaving(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"savings": savings})
}

func UpdateSaving(c *gin.Context) {
	var input SavingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := models.Saving{}
	v.ID = uint(id)
	v.UserID = userId
	v.Amount = input.Amount

	_, err = v.UpdateSaving()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Saving updated."})
}

func DeleteSaving(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := models.Saving{}
	v.ID = uint(id)
	v.UserID = userId

	_, err = v.DeleteSaving()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Saving deleted."})
}

func GetSavingRisk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"risk": 0})
}

func GetSavingReturn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"return": 3})
}
