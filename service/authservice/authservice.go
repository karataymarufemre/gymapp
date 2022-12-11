package authservice

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GenerateJWT(userId uint64, minutes time.Duration, isLongToken bool) (string, error)
	VerifyJWT(tokenString string) (JWTClaims, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}

type AuthServiceImpl struct {
	JwtKey []byte
}

type JWTClaims struct {
	jwt.RegisteredClaims
	UserId      uint64 `json:"user_id"`
	IsLongToken bool   `json:"is_long_token"`
}

func NewAuthService(JwtKey []byte) *AuthServiceImpl {
	return &AuthServiceImpl{
		JwtKey: JwtKey,
	}
}

func (a *AuthServiceImpl) GenerateJWT(userId uint64, minutes time.Duration, isLongToken bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(minutes * time.Minute))},
		UserId:           userId,
		IsLongToken:      isLongToken,
	})
	tokenString, err := token.SignedString(a.JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *AuthServiceImpl) VerifyJWT(tokenString string) (JWTClaims, error) {
	claims := JWTClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.JwtKey, nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil
}

func (a *AuthServiceImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (a *AuthServiceImpl) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
