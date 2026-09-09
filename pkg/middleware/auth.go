package middleware

import (
	"net/http"
	"strings"
	"temporary/pkg/jwt"
)

type AuthMiddleware struct {
	JWT *jwt.JWT
}

func (a *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(
			r.Header.Get("Authorization"),
			"Bearer ",
		)

		if token == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
		}

		if _, err := a.JWT.Verify(token); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
		})
}