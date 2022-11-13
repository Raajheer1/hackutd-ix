package controllers

import (
	"github.com/Raajheer1/hackutd-ix/m/v2/models"
	"github.com/Raajheer1/hackutd-ix/m/v2/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BondInput struct {
	UserID   uint   `json:"user_id" binding:"required"`
	Symbol   string `json:"symbol" binding:"required"`
	Quantity string `json:"quantity" binding:"required"`
}

func CreateBond(c *gin.Context) {
	var input BondInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := models.Bond{}
	b.UserID = input.UserID
	b.Symbol = input.Symbol
	b.Quantity = input.Quantity

	_, err := b.CreateBond()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bond Added"})
}

func GetBond(c *gin.Context) {
	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bonds, err := models.GetBond(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bonds": bonds})
}

func UpdateBond(c *gin.Context) {
	var input BondInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := models.Bond{}
	b.ID = uint(id)
	b.UserID = input.UserID
	b.Symbol = input.Symbol
	b.Quantity = input.Quantity

	_, err = b.UpdateBond()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bond updated."})
}

func DeleteBond(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := models.Bond{}
	b.ID = uint(id)

	_, err = b.DeleteBond()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bond deleted."})
}
