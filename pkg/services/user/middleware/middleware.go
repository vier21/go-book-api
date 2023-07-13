package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vier21/go-book-api/config"
	"github.com/vier21/go-book-api/pkg/services/user/service"
)

func TimoutMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func VerifyJWT(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		raw := strings.Split(r.Header.Get("Authorization"), " ")

		if len(raw) < 2 {
			http.Error(w, "you're not authorized", http.StatusUnauthorized)
			log.Println("Invalid token")
			return
		} else if len(raw) > 2 {
			http.Error(w, "you're not authorized", http.StatusUnauthorized)
			log.Println("Invalid token")
			return
		}

		bearer := raw[1]

		if bearer == "" {
			http.Error(w, "you're not authorized", http.StatusUnauthorized)
			log.Println("Invalid token")
			return
		}

		var claims service.JWTClaims

		token, err := jwt.ParseWithClaims(bearer, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.GetConfig().SecretKey), nil
		})

		var ctx context.Context
		if claim, ok := token.Claims.(*service.JWTClaims); ok && token.Valid {
			ctx = context.WithValue(r.Context(), "data", claim.Data)
		} else {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			log.Printf("Error %s", err)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
