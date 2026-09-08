package api

import (
	"Saavedra/service/Quote/service"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type EndpointHandler struct {
	service service.Service
}

func New(service service.Service) EndpointHandler {
	return EndpointHandler{service: service}
}

// ROUTE: /quote/menu
func CallQuoteMenu(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tpl, err := template.ParseFiles("service/Quote/views/menu.html")
		if err != nil {
			log.Printf("Error parsing HTML /quote/menu...(%v)", err.Error())
			http.Error(w, "Error parsing HTML /quote/menu", http.StatusInternalServerError)
			return
		}
		if err = tpl.Execute(w, nil); err != nil {
			log.Printf("Error rendering HTML /quote/menu...(%v)", err.Error())
			http.Error(w, "Error rendering HTML /quote/menu", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// ROUTE: /quote/new
func CallQuoteNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tpl, err := template.ParseFiles("service/Quote/views/quoteNew.html")
		if err != nil {
			log.Printf("Error parsing HTML /quote/new...(%v)", err.Error())
			http.Error(w, "Error parsing HTML /quote/new", http.StatusInternalServerError)
			return
		}
		if err = tpl.Execute(w, nil); err != nil {
			log.Printf("Error rendering HTML /quote/new...(%v)", err.Error())
			http.Error(w, "Error rendering HTML /quote/new", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// ROUTE: /quote/many2one
func (e EndpointHandler) CallMany2One(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		many2one, err := e.service.Many2One()
		if err != nil {
			log.Printf("Error fetching many2one data /quote/many2one...(%v)", err.Error())
			http.Error(w, "Error fetching many2one data /quote/many2one", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(many2one)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
