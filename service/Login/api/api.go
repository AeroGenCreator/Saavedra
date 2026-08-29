package api

import (
	"Saavedra/env"
	"Saavedra/service/Login/service"
	"Saavedra/service/Login/types"
	"Saavedra/utils"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type EndpointHandler struct {
	service service.Service
}

func New(s service.Service) *EndpointHandler {
	return &EndpointHandler{service: s}
}

// SERVES -> /login
func (e *EndpointHandler) CallLogin(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		brand, err := e.service.LoadCompanyBranding()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl, err := template.ParseFiles("service/Login/views/login.html")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, brand)

	case http.MethodPost:
		var request *types.User

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		creds, err := e.service.ValidatePassword(request)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if !creds.ValidationStatus {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    creds.Token,
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
			HttpOnly: true,
			Secure:   env.IsProduction,
			SameSite: http.SameSiteLaxMode,
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"success":  true,
			"username": creds.UserName,
		})

	case http.MethodPatch:

		// WHENEVER PUT METHOD IS CALL -> DB APY TOKEN IS SET TO ''.
		userID, ok := r.Context().Value(utils.UserIDKey).(string)

		if !ok {
			log.Println("Cookies extraction failed /login...")
			http.Error(w, "/login - Method PUT", http.StatusInternalServerError)
			return
		}

		intID, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing 'string'", http.StatusInternalServerError)
			return
		}

		authUser := types.AuthUserSession{
			UserId:    intID,
			AuthToken: "",
		}

		if err = e.service.LogOut(&authUser); err != nil {
			log.Println(err)
			http.Error(w, "DB Alteration Error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    "",
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
			HttpOnly: true,
			Secure:   env.IsProduction,
			SameSite: http.SameSiteLaxMode,
		})

		w.WriteHeader(http.StatusOK)
		return

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}

// SERVES -> /
func (e *EndpointHandler) CallRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}

func (e *EndpointHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		apiKey, ok1 := r.Context().Value(utils.ApiKey).(string)
		userId, ok2 := r.Context().Value(utils.UserIDKey).(string)
		userName, ok3 := r.Context().Value(utils.UserNameKey).(string)

		if !ok1 || !ok2 || !ok3 {
			log.Println("Cookies extraction failed /refresh...")
			http.Error(w, "/login - Method PUT", http.StatusInternalServerError)
			return
		}

		intID, err := strconv.Atoi(userId)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error parsing 'string'", http.StatusInternalServerError)
			return
		}

		log.Println("Refresh token attempting...")

		authUser := &types.AuthUserSession{
			UserId:    intID,
			AuthToken: apiKey,
		}

		newToken, err := e.service.ValidateTokenRefresh(userName, authUser)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if newToken == "" {
			e.service.LogOut(&types.AuthUserSession{
				UserId:    authUser.UserId,
				AuthToken: "",
			})
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// COOKIES CREATION

		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    newToken,
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
			HttpOnly: true,
			Secure:   env.IsProduction,
			SameSite: http.SameSiteLaxMode,
		})

		w.WriteHeader(http.StatusOK)
		return

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
