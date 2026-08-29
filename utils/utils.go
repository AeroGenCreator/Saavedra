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

const (
	UserIDKey   contextKey = "userID"
	ApiKey      contextKey = "apiKEY"
	UserNameKey contextKey = "userNAME"
)

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
func GenerateToken(name string, id int) (string, error) {
	var jwtKey = []byte(env.ApyKey)
	expirationTime := time.Now().Add(30 * time.Second)

	claims := &types.Claims{
		UserId:   strconv.Itoa(id),
		UserName: name,
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

		// JWTKEY
		jwtKey := []byte(os.Getenv("JWT_KEY"))
		requestedWith := r.Header.Get("X-Requested-With")
		log.Println(requestedWith)
		// APIKEY FROM JS
		c, err := r.Cookie("auth_token")

		if err == http.ErrNoCookie {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// TOKEN EXTRACTION && VALIDATION
		tokenStr := c.Value
		claims := &types.Claims{}

		_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			return jwtKey, nil
		})

		if err != nil {
			// HANDLES INVALID TOKEN OR NIL -> FROM JS RETURNS 401 -> FROM URL RETURN /login
			log.Println(err.Error())

			if requestedWith != "jsFrontendComponent" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// USER ID IS ADDED TO THE REQUEST
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserId)
		ctx = context.WithValue(ctx, UserNameKey, claims.UserName)
		ctx = context.WithValue(ctx, ApiKey, tokenStr)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func LowLevelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("auth_token")

		// 1. Manejo de errores de la cookie
		if err != nil {
			if err == http.ErrNoCookie {
				next.ServeHTTP(w, r)
				return
			}

			log.Println("Cookies error...", err)
			http.Error(w, "Invalid Cookies", http.StatusBadRequest)
			return
		}

		tokenStr := c.Value

		// 2. Si la cookie está vacía
		if tokenStr == "" {
			next.ServeHTTP(w, r)
			return
		}

		// 3. Parsear token
		claims := &types.Claims{}
		parser := jwt.NewParser()

		_, _, err = parser.ParseUnverified(tokenStr, claims)
		if err != nil {
			log.Println("Error al parsear token no verificado:", err)
			next.ServeHTTP(w, r)
			return
		}

		// 4. Inyección exitosa del valor en el contexto
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserId)
		ctx = context.WithValue(ctx, UserNameKey, claims.UserName)
		ctx = context.WithValue(ctx, ApiKey, tokenStr)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// PARSE STRING TO NULL IF EMPTY
func StringONull(input string) *string {
	if input == "" {
		return nil
	}

	return &input
}
