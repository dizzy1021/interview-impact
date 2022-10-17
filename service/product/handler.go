package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service *ProductService) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Product Added")
	}	
}

func (service *ProductService) RemoveProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Product Removed")
	}	
}

func (service *ProductService) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Product Updated")
	}	
}

func (service *ProductService) FindProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(http.StatusOK, "Product Found: " + id)
	}	
}

func (service *ProductService) FindAllProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Products Found")
	}	
}