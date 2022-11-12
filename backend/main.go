package main

import (
	"fmt"
	"github.com/Raajheer1/hackutd-ix/m/v2/controllers"
	"github.com/Raajheer1/hackutd-ix/m/v2/middlewares"
	"github.com/Raajheer1/hackutd-ix/m/v2/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading the environment file.")
		return
	}

	models.ConnectDatabase()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{os.Getenv("APP_URL")}
	corsConfig.AddAllowMethods("OPTIONS")

	r := gin.Default()
	r.Use(cors.New(corsConfig))

	auth := r.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.GetCurrentUser)

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		fmt.Println("Error occurred during server initialization.")
		return
	}
}
