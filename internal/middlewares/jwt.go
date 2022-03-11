package middlewares

import (
	"net/http"
	"strings"

	"github.com/blyndusk/flamingops/internal/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func JWTVerify(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")[1]

		jwtWrapper := utils.JwtWrapper{
			SecretKey: utils.GetJWTSecretKey(),
			Issuer:    "AuthService",
		}

		Claims, err := jwtWrapper.ValidateToken(jwtString)
		log.Error(err)

		if err != nil {
			c.JSON(http.StatusUnauthorized, "token not valid")
			return
		}

		c.Set("userId", Claims.UserID)
		next(c)
	}
}
