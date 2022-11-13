package main

import (
	"fmt"
	"github.com/Raajheer1/hackutd-ix/m/v2/controllers"
	"github.com/Raajheer1/hackutd-ix/m/v2/middlewares"
	"github.com/Raajheer1/hackutd-ix/m/v2/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Risk Management API Online."})
	})

	auth := r.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	//Users
	protected.GET("/user", controllers.GetCurrentUser)

	//Stocks
	stock := protected.Group("/stock")
	stock.GET("/", controllers.GetStock)
	stock.POST("/", controllers.CreateStock)
	stock.PUT("/:id", controllers.UpdateStock)
	stock.DELETE("/:id", controllers.DeleteStock)
	stock.GET("/return", controllers.GetStockReturn)
	stock.GET("/risk", controllers.GetStockRisk)

	//Savings
	saving := protected.Group("/saving")
	saving.GET("/", controllers.GetSaving)
	saving.POST("/", controllers.CreateSaving)
	saving.PUT("/:id", controllers.UpdateSaving)
	saving.DELETE("/:id", controllers.DeleteSaving)
	saving.GET("/return", controllers.GetSavingReturn)
	saving.GET("/risk", controllers.GetSavingRisk)

	//Bonds
	bond := protected.Group("/bond")
	bond.GET("/", controllers.GetBond)
	bond.POST("/", controllers.CreateBond)
	bond.PUT("/:id", controllers.UpdateBond)
	bond.DELETE("/:id", controllers.DeleteBond)
	bond.GET("/return", controllers.GetBondReturn)

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		fmt.Println("Error occurred during server initialization.")
		return
	}
}
