package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Raajheer1/hackutd-ix/m/v2/models"
	"github.com/Raajheer1/hackutd-ix/m/v2/utils/token"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
)

type BondInput struct {
	UserID       uint   `json:"user_id" binding:"required"`
	PurchaseDate string `json:"purchase_date" binding:"required"`
	Maturity     string `json:"maturity" binding:"required"`
}

func CreateBond(c *gin.Context) {
	var input BondInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := models.Bond{}
	b.UserID = input.UserID
	b.PurchaseDate = input.PurchaseDate
	b.Maturity = input.Maturity

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

	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := models.Bond{}
	b.ID = uint(id)
	b.UserID = userId
	b.PurchaseDate = input.PurchaseDate
	b.Maturity = input.Maturity

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

	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := models.Bond{}
	b.ID = uint(id)
	b.UserID = userId

	_, err = b.DeleteBond()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bond deleted."})
}

type BondsAPIResponse struct {
	Data []BondPrice `json:"data"`
}
type BondPrice struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

func GetBondReturn(c *gin.Context) {
	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bonds, err := models.GetBond(userId)

	percentage := 0.0
	count := 0.0
	for _, bond := range bonds {
		resp, err := http.Get(fmt.Sprintf("https://www.alphavantage.co/query?function=TREASURY_YIELD&interval=monthly&maturity=%s&apikey=%s", bond.Maturity, os.Getenv("API_KEY")))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		var result BondsAPIResponse
		if err := json.Unmarshal(body, &result); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for _, day := range result.Data {
			if day.Date == bond.PurchaseDate {
				count += 1.0
				if rate, err := strconv.ParseFloat(day.Value, 32); err == nil {
					percentage += rate
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
			}
		}
	}

	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No valid bonds found."})
	} else {
		c.JSON(http.StatusOK, gin.H{"return": percentage / count})
	}
	return
}
