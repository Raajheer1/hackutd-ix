package controllers

import (
	"github.com/Raajheer1/hackutd-ix/m/v2/models"
	"github.com/Raajheer1/hackutd-ix/m/v2/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCurrentUser(c *gin.Context) {
	userId, err := token.ExtractID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password

	authToken, err := models.Login(u.Email, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": authToken})
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Income   uint   `json:"income" binding:"required"`
	Age      uint   `json:"age" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password
	u.Income = input.Income
	u.Age = input.Age

	_, err := u.CreateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}
