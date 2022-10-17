package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router  := gin.Default()

	port := loadEnv("PORT")
	readTimeout, err := strconv.Atoi(loadEnv("READ_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error load read timeout env")
	}

	writeTimeout, err := strconv.Atoi(loadEnv("WRITE_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error load write timeout env")
	}

	appMode := loadEnv("APP_MODE")
	appName := loadEnv("APP_NAME")
	
	server := &http.Server{
		Addr		: ":" + port,
		Handler		: router,
		ReadTimeout	: time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,		
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": appName + " API",
		})
	})

	gin.SetMode(appMode)

	server.ListenAndServe()
}

func loadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}