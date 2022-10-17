package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductService struct {
	engine *gin.Engine
	store *ProductStore
}

func NewProductService(engine *gin.Engine, db *gorm.DB) *ProductService {
	return &ProductService{
		engine: engine,
		store: NewProductStore(db),
	}
}

func (service *ProductService) New(routerGroup *gin.RouterGroup) {
	router := service.engine.Group("/")
	if router != nil {
		router = routerGroup
	}	
	router.POST("/product", service.CreateProduct())
	router.PUT("/product", service.UpdateProduct())
	router.DELETE("/product", service.RemoveProduct())
	router.GET("/product", service.FindAllProduct())
	router.GET("/product/:id", service.FindProduct())
}