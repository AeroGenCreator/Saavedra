package api

import (
	"Saavedra/env"
	"Saavedra/service/Login/service"
	"Saavedra/service/Login/types"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
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
			Expires:  time.Now().Add(5 * time.Minute),
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
