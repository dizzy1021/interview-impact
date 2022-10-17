package middleware

import (	
	"net/http"

	"dizzy1021.dev/interview-impact/util"
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, password, ok := ctx.Request.BasicAuth()
		if !ok {
			message := "autentikasi tidak ditemukan"
			resp := util.NewAPIResponse(nil, &message, http.StatusUnauthorized)			
			ctx.JSON(http.StatusUnauthorized, resp)
			ctx.Abort()
			return
		}

		validUsername := util.LoadEnv("BASIC_AUTH_USERNAME")
		validPassword := util.LoadEnv("BASIC_AUTH_PASSWORD")

		isValid := (username == validUsername) && (password == validPassword)
		if !isValid {
			message := "autentikasi tidak valid"			
			resp := util.NewAPIResponse(nil, &message, http.StatusUnauthorized)
			ctx.JSON(http.StatusUnauthorized, resp)
			ctx.Abort()
			return
		}

		ctx.Next()
	}	
}
