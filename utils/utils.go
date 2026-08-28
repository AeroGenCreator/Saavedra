package utils

import (
	"Saavedra/env"
	"Saavedra/service/Login/types"
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// === REUSABLES ===

type contextKey string

const UserIDKey contextKey = "userID"

// FUNCTION -> HASHED STRINGS
func HashPassword(password string) (string, error) {
	// GenerateFromPassword expects a byte slice and a cost factor (e.g., bcrypt.DefaultCost)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FUNCTION: CHECK PASSWORD
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// FUNCTION: GENERATE SIGNED TOKEN
func GenerateToken(user *types.User) (string, error) {
	var jwtKey = []byte(env.ApyKey)
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &types.Claims{
		UserId:   strconv.Itoa(user.Id),
		UserName: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// JWT - VALIDATE TOKEN => (Name: "auth_token")
// : Must protect all endpoints.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TOKEN && API KEY
		c, err := r.Cookie("auth_token")
		jwtKey := []byte(os.Getenv("JWT_KEY"))

		if err != nil {
			if err == http.ErrNoCookie {
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// TOKEN EXTRACTION && VALIDATION
		tokenStr := c.Value
		claims := &types.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if err != nil || !tkn.Valid {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// USER ID IS ADDED TO THE REQUEST
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserId)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
