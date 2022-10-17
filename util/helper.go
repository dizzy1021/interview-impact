package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func LoadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page == 0 { 
			page = 1
		}

		switch {
		case pageSize > MAX_PAGE_SIZE:
		  pageSize = MAX_PAGE_SIZE
		case pageSize <= 0:
		  pageSize = MIN_PAGE_SIZE
		}
	
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}