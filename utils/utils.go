package utils

import (
	"Saavedra/service/Login/types"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// === REUSABLES ===

// FUNCTION -> HASHED STRINGS
func HashPassword(password string) (string, error) {
	// GenerateFromPassword expects a byte slice and a cost factor (e.g., bcrypt.DefaultCost)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FUNCTION: GENERATE SIGNED TOKEN
func GenerateToken(userName string) (string, error) {
	Key := string(os.Getenv("JWT_KEY"))
	var jwtKey = []byte(Key)
	if Key == "" {
		return "", errors.New("System cannot continue since environment 'JWT_KEY' has not been declared")
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &types.Claims{
		UserName: userName,
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

		next.ServeHTTP(w, r)

	})
}
