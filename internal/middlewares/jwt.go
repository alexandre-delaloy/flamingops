package middlewares

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/blyndusk/flamingops/internal/utils"
)

func JWTVerify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtWrapper := utils.JwtWrapper{
			SecretKey: utils.GetJWTSecretKey(),
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(r.Header.Get("Authorization"))

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("token not valid")
			return
		} else {
			r.Header.Set("user_id", strconv.Itoa(int(claims.UserID)))
			next(w, r)
		}
	}
}
