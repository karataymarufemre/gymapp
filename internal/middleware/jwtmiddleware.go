package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"workspace/internal/constants/contextkey"
	"workspace/internal/constants/urlconstants"
	"workspace/internal/service/authservice"
	"workspace/internal/utils/sliceutils"
)

type JWTMiddleware interface {
	Middleware(next http.Handler) http.Handler
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
		if sliceutils.Contains(urlconstants.NO_JWT(), r.RequestURI) {
			next.ServeHTTP(w, r)
		} else if sliceutils.Contains(urlconstants.AUTH_ALLOWED_REFRESH_JWT(), r.RequestURI) {
			j.middlewareForAccessToken(next).ServeHTTP(w, r)
		} else {
			j.middlewareForNormal(next).ServeHTTP(w, r)
		}
	})
}

func (j *JWTMiddlewareImpl) middlewareForNormal(next http.Handler) http.Handler {
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
				ctx := context.WithValue(r.Context(), contextkey.CLAIMS_KEY, claims)
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

func (j *JWTMiddlewareImpl) middlewareForAccessToken(next http.Handler) http.Handler {
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
				ctx := context.WithValue(r.Context(), contextkey.CLAIMS_KEY, claims)
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
