package middleware

import (
	"chuxin0816/SE/common"
	"chuxin0816/SE/models"
	"chuxin0816/SE/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		userId := claims.UserId
		var user models.User
		common.DB.First(&user, userId)
		if user.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
