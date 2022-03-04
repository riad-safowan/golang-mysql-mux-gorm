package middleware

import (
	"github.com/gorilla/context"
	"net/http"
	"strings"

	"github.com/riad-safowan/GOLang-SQL/helpers"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		clientToken := r.Header.Get("Authorization")
		if clientToken == "" {
			clientToken = r.Header.Get("token")
		} else if strings.HasPrefix(clientToken, "Bearer ") {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			clientToken = splitToken[1]
		} else {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			return
		}

		if clientToken == "" {
			http.Error(w, "No Authorization header provided", http.StatusUnauthorized)
			return
		}
		// handle access token
		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			http.Error(w, err, http.StatusUnauthorized)
			return
		}

		if claims.Token_type == "access_token" {
			context.Set(r, "email", claims.Email)
			context.Set(r, "user_id", claims.Uid)
		} else if claims.Token_type == "refresh_token" {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
		}

		next(w, r)
	}
}
