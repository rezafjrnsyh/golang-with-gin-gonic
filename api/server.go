package api

import (
	"baf/api/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)
func Run() {

	gin.SetMode(gin.ReleaseMode)

	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	db := config.ConnectDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	r := config.CreateRouter()
	config.InitRouter(db, r).InitializeRoutes()
	config.Run(r, ":8800")

}
