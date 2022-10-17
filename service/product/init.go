package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductService struct {
	router *gin.Engine
	store *ProductStore
}

func NewProductService(router *gin.Engine, db *gorm.DB) *ProductService {
	return &ProductService{
		router: router,
		store: NewProductStore(db),
	}
}