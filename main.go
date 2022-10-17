package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"dizzy1021.dev/interview-impact/driver"
	"dizzy1021.dev/interview-impact/middleware"
	"dizzy1021.dev/interview-impact/model"
	"dizzy1021.dev/interview-impact/service/product"
	"dizzy1021.dev/interview-impact/util"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {
	router  := gin.Default()
	router.ForwardedByClientIP = true

	port := util.LoadEnv("PORT")
	readTimeout, err := strconv.Atoi(util.LoadEnv("READ_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error load read timeout env")
	}

	db := driver.LoadDatabase()

	writeTimeout, err := strconv.Atoi(util.LoadEnv("WRITE_TIMEOUT"))
	if err != nil {
		log.Fatalf("Error load write timeout env")
	}

	appMode := util.LoadEnv("APP_MODE")
	appName := util.LoadEnv("APP_NAME")
	
	server := &http.Server{
		Addr		: ":" + port,
		Handler		: router,
		ReadTimeout	: time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,			
	}

	gin.SetMode(appMode)	

	// DB migration
	migrate(db)

	// Root Endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": appName + " API",
		})
	})

	// Register Services
	v1 := router.Group("/api/v1", middleware.BasicAuth(), middleware.RateLimitter())
	product.NewProductService(router, db).New(v1)

	server.ListenAndServe()
}

func migrate(db *gorm.DB) {
	// Create enum uom
	query := fmt.Sprintf(`DROP TYPE IF EXISTS uom CASCADE; CREATE TYPE uom AS ENUM ('%s', '%s', '%s')`, model.SHEET, model.ROLL, model.PCS)
	db.Exec(query)
	
	// Auto Migrate
	db.AutoMigrate(&(model.Product{}))
}