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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": message})
			return
		}

		validUsername := util.LoadEnv("BASIC_AUTH_USERNAME")
		validPassword := util.LoadEnv("BASIC_AUTH_PASSWORD")

		isValid := (username == validUsername) && (password == validPassword)
		if !isValid {        
			message := "autentikasi tidak valid"
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": message})
			return
		}

		ctx.Next()
	}	
}
