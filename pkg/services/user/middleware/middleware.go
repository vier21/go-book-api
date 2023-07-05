package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vier21/go-book-api/config"
	"github.com/vier21/go-book-api/pkg/services/user/service"
)

func TimoutMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		next(w, r.WithContext(ctx))
	}
}

func VerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearer := strings.Split(r.Header.Get("Authorization"), " ")
		
		if len(bearer) > 2 {
			http.Error(w, "invalid Token", http.StatusUnauthorized)
		}

		var claims service.JWTClaims

		token, err := jwt.ParseWithClaims(bearer[1], &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.GetConfig().SecretKey), nil
		})

		if claim, ok := token.Claims.(*service.JWTClaims); ok && token.Valid {
			fmt.Printf("%v %v", claim.Data, claim.RegisteredClaims.Issuer)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Printf("Error %s", err)
			return
		}

		next.ServeHTTP(w, r)

	}
}
