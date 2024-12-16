package jwtauthmiddleware

import (
	"link-back-app/api"
	tokenutil "link-back-app/utils/token_util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := tokenutil.TokenValid(context)
		if err != nil {
			api.ErrorrResponse(context, http.StatusUnauthorized, "トークンの有効期限が切れています")
			context.Abort()
			return
		}
		context.Next()
	}
}
