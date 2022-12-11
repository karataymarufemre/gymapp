package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"workspace/service/authservice"
)

type JWTMiddleware interface {
	Middleware(next http.Handler) http.Handler
	MiddlewareForAccessToken(next http.Handler) http.Handler
}

type JWTMiddlewareImpl struct {
	authService authservice.AuthService
}

func NewJWTMiddleware(authService authservice.AuthService) *JWTMiddlewareImpl {
	return &JWTMiddlewareImpl{
		authService: authService,
	}
}

func (j *JWTMiddlewareImpl) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			claims, err := j.authService.VerifyJWT(jwtToken)
			if err == nil && !claims.IsLongToken {
				ctx := context.WithValue(r.Context(), "claims", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

func (j *JWTMiddlewareImpl) MiddlewareForAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			claims, err := j.authService.VerifyJWT(jwtToken)
			if err == nil && claims.IsLongToken {
				ctx := context.WithValue(r.Context(), "claims", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}
