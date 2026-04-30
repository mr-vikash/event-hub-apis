package middleware

import (
	"context"
	"net/http"
	"strings"

	"eventhub/utils"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

var UserContextKey = contextKey("user_id")

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Split(authHeader, "Bearer ")
		if len(tokenStr) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenStr[1], func(token *jwt.Token) (interface{}, error) {
			return utils.SECRET, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		ctx := context.WithValue(r.Context(), UserContextKey, userID)
		next(w, r.WithContext(ctx))
	}
}
