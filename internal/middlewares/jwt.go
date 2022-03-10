package middlewares

import (
	"net/http"
	"strings"
	"strconv"

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

		claims, err := jwtWrapper.ValidateToken(jwtString)
		log.Error(err)

		if err != nil {
			c.JSON(http.StatusUnauthorized, "token not valid")
			return
		}


		i64, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "mistakes were made")
			return
		}

		if claims.UserID != uint(i64) {
			c.JSON(http.StatusUnauthorized, "wrong id")
			return
		}

		next(c)
	}
}