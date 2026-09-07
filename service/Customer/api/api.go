package api

import (
	"Saavedra/service/Customer/service"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type EndpointHandler struct {
	service service.Service
}

func New(servie service.Service) *EndpointHandler {
	return &EndpointHandler{service: servie}
}

// ROUTE: /customer
func (e EndpointHandler) CallCustomer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tpl, err := template.ParseFiles("service/Customer/views/customer.html")
		if err != nil {
			log.Printf("Error parsing /customer...(%v)", err.Error())
			http.Error(w, "Error parsing /customer", http.StatusInternalServerError)
			return
		}
		if err = tpl.Execute(w, nil); err != nil {
			log.Printf("Error rendering /customer...(%v)", err.Error())
			http.Error(w, "Error rendering /customer", http.StatusInternalServerError)
		}
	case http.MethodPut:
		return
	case http.MethodPost:
		return
	case http.MethodDelete:
		return
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// ROUTE: /customer/slice
func (e EndpointHandler) CallCustomerList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		page := r.URL.Query().Get("page")
		records, err := e.service.ListCustomer(page)
		if err != nil {
			log.Printf("Error fetchin data /customer/slice...(%v)", err.Error())
			http.Error(w, "Error fetchin data /customer/slice", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(records); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Error internal server", http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		return
	case http.MethodPost:
		return
	case http.MethodDelete:
		return
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// ROUTE: /customer/new
func (e EndpointHandler) CallCustomerNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPost:
		return
	case http.MethodDelete:
		return
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}

// ROUTE: /customer/record
func (e EndpointHandler) CallCustomerRecord(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPost:
		return
	case http.MethodDelete:
		return
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
