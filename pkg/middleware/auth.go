package middleware

import (
	"net/http"
	"temporary/pkg/jwt"
	"strings"
)

func AuthMiddleware(jwtService *jwt.JWT) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := strings.TrimPrefix(
                r.Header.Get("Authorization"),
                "Bearer ",
            )

            if token == "" {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            if _, err := jwtService.Verify(token); err != nil {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}