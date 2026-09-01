package api

import (
	"Saavedra/service/Users/service"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type EndpointHandler struct {
	service service.Service
}

func New(s service.Service) *EndpointHandler {
	return &EndpointHandler{service: s}
}

// ROUTE: /users -> ONLY RENDERS HTML
func (e EndpointHandler) CallUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		tpl, err := template.ParseFiles("service/Users/views/users.html")
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Couldn't parse users HTML template file.", http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, nil)
		if err != nil {
			log.Print("Error when rendering users template.")
			http.Error(w, "Render Error", http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}

// ROUTE: /users/records?page=x -> ONLY RETURNS TYPE(TABLE) -> RECORDS & METADATA
func (e EndpointHandler) CallUsersRecords(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		page := r.URL.Query().Get("page")
		if page != "" {
			page = "1"
		}

		intPage, err := strconv.Atoi(page)
		if err != nil {
			intPage = 1
		}

		data, err := e.service.ListUsers(intPage)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Fetching users data error...", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return

	}

}

// ROUTE: /users/new -> REDIRECTS -> /users/form?id=x
func (e EndpointHandler) CreateUserInDB(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		tpl, err := template.ParseFiles("service/Users/views/usersCreate.html")

		if err != nil {
			log.Print("Error parsing usersCreate HTML template...")
			http.Error(w, "Parsing error...", http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, nil)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Loading error...", http.StatusInternalServerError)
			return
		}

	case http.MethodHead:
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}
