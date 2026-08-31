package api

import (
	"Saavedra/service/Users/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type EndpointHandler struct {
	service service.Service
}

func New(s service.Service) *EndpointHandler {
	return &EndpointHandler{service: s}
}

func (e EndpointHandler) CallUsersByPages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		pageStr := strings.TrimPrefix(r.URL.Path, "/users/")
		page, err := strconv.Atoi(pageStr)

		if err != nil {
			page = 1
		}

		table, err := e.service.ListUsers(page)

		tpl, err := template.ParseFiles("service/Users/views/users.html")
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Couldn't parse users HTML template file.", http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, table)
		if err != nil {
			log.Print("Error when rendering users template.")
			http.Error(w, "Render Error", http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}
