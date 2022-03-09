package middlewares

import (
	"net/http"

	"github.com/blyndusk/flamingops/internal/utils"
	"github.com/gin-gonic/gin"
)

func JWTVerify(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtWrapper := utils.JwtWrapper{
			SecretKey: utils.GetJWTSecretKey(),
			Issuer:    "AuthService",
		}

		_, err := jwtWrapper.ValidateToken(c.Request.Header.Get("Authorization"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, "token not valid")
			return
		}

		next(c)
	}
}
