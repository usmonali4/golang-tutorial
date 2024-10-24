package main

import (
	// "net/http"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/usmonali4/jwt_auth/controllers"
	"github.com/usmonali4/jwt_auth/models"
)

func main() {

	err := godotenv.Load("../.env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	models.ConnectDB()

	r := gin.Default()
	portStr := os.Getenv("PORT")

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	r.Run(":" + portStr)
}