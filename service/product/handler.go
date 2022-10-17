package product

import (	
	"net/http"

	"dizzy1021.dev/interview-impact/model"
	"dizzy1021.dev/interview-impact/util"
	"github.com/gin-gonic/gin"
)

func (service *ProductService) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product model.Product

		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			errors := err.Error()
			resp := util.NewAPIResponse(nil, errors, http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, resp)
			ctx.Abort()
			return
		}

		// Validate Unique Code
		exists := service.store.FindOneProductByCode(product.Code)		
		if exists.ID != "" {
			message := "kode produk sudah tersedia"
			resp := util.NewAPIResponse(nil, message, http.StatusInternalServerError)
			ctx.JSON(http.StatusBadRequest, resp)
			ctx.Abort()
			return
		}		

		err = service.store.InsertProduct(product)
		if err != nil {
			message := "gagal menambahkan produk " + err.Error()
			resp := util.NewAPIResponse(nil, message, http.StatusInternalServerError)			
			ctx.JSON(http.StatusInternalServerError, resp)
			ctx.Abort()
			return
		}

		message := "berhasil menambahkan produk"
		resp := util.NewAPIResponse(nil, message, http.StatusOK)
		ctx.JSON(http.StatusOK, resp)
	}	
}

func (service *ProductService) RemoveProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := service.store.DeleteProduct(id)
		if err != nil {
			message := "gagal menghapus produk " + err.Error()
			resp := util.NewAPIResponse(nil, message, http.StatusInternalServerError)			
			ctx.JSON(http.StatusInternalServerError, resp)
			ctx.Abort()
			return
		}
		
		message := "berhasil menghapus produk"
		resp := util.NewAPIResponse(nil, message, http.StatusOK)
		ctx.JSON(http.StatusOK, resp)
	}	
}

func (service *ProductService) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		id := ctx.Param("id")		
		var product model.Product

		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			errors := err.Error()
			resp := util.NewAPIResponse(nil, errors, http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, resp)
			ctx.Abort()
			return
		}

		// Validate product exist
		_, err = service.store.FindOneProduct(id)
		if err != nil {
			message := "produk tidak ditemukan"
			resp := util.NewAPIResponse(nil, message, http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, resp)
			ctx.Abort()
			return
		}

		product.ID = id
		err = service.store.UpdateProduct(product)
		if err != nil {
			message := "gagal mengupdate produk " + err.Error()
			resp := util.NewAPIResponse(nil, message, http.StatusInternalServerError)			
			ctx.JSON(http.StatusInternalServerError, resp)
			ctx.Abort()
			return
		}

		message := "berhasil mengupdate produk"
		resp := util.NewAPIResponse(nil, message, http.StatusOK)
		ctx.JSON(http.StatusOK, resp)
	}	
}

func (service *ProductService) FindProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
				
		product, err := service.store.FindOneProduct(id)
		if err != nil {
			message := "produk tidak ditemukan"
			resp := util.NewAPIResponse(nil, message, http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, resp)
			ctx.Abort()
			return
		}

		message := "produk ditemukan"
		resp := util.NewAPIResponse(product, message, http.StatusOK)
		ctx.JSON(http.StatusOK, resp)
	}	
}

func (service *ProductService) FindAllProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Products Found")
	}	
}