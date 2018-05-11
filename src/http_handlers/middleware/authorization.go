package middleware

import (
	"net/http"
	"github.com/Maykonn/jwt-go-validation"
	"os"
	"fmt"
)

type Authorization struct{}

func (auth *Authorization) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tokenStr := r.Header.Get("Authorization"); tokenStr != "" {
			_, err := jwt_go_validation.JwtParse(os.Getenv("JWT_KEY"), tokenStr)
			if err == nil {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
				fmt.Println(err.Error())
			}
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println("Authorization token is required")
		}
	})
}
