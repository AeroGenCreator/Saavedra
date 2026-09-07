package api

import (
	"Saavedra/service/Customer/service"
	"Saavedra/service/Customer/types"
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
		tpl, err := template.ParseFiles("service/Customer/views/customerNew.html")
		if err != nil {
			log.Printf("Error parsing HTML /customer/new...(%v)", err.Error())
			http.Error(w, "Error parsing HTML /customer/new", http.StatusInternalServerError)
			return
		}
		if err = tpl.Execute(w, nil); err != nil {
			log.Printf("Error rendering HTML /customer/new...(%v)", err.Error())
			http.Error(w, "Error rendering HTML /customer/new", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var customer types.Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			log.Printf("Error decoding JSON /customer/new...(%v)", err.Error())
			http.Error(w, "Error decoding JSON /customer/new", http.StatusInternalServerError)
			return
		}
		if err := e.service.CreateCustomer(&customer); err != nil {
			log.Printf("Error creating /customer/new...(%v)", err.Error())
			http.Error(w, "Error creating /customer/new", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
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
		id := r.URL.Query().Get("id")
		record, err := e.service.ReadCustomer(id)
		if err != nil {
			log.Printf("Error query /customer/record...(%v)", err.Error())
			http.Error(w, "Error query /customer/record", http.StatusInternalServerError)
			return
		}
		recordBytes, err := json.Marshal(record)
		if err != nil {
			log.Printf("Error parsing bytes /customer/record...(%v)", err.Error())
			http.Error(w, "Error parsing bytes /customer/record", http.StatusInternalServerError)
			return
		}
		tpl, err := template.ParseFiles("service/Customer/views/customerRecord.html")
		if err != nil {
			log.Printf("Error parsing template /customer/record...(%v)", err.Error())
			http.Error(w, "Error parsing template /customer/record", http.StatusInternalServerError)
			return
		}
		if err = tpl.Execute(w, map[string]template.JS{"Record": template.JS(recordBytes)}); err != nil {
			log.Printf("Error rendering template /customer/record...(%v)", err.Error())
			http.Error(w, "Error rendering template /customer/record", http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		var customer types.Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			log.Printf("Error decoding PUT /customer/record...(%v)", err.Error())
			http.Error(w, "Error decoding PUT /customer/record", http.StatusInternalServerError)
			return
		}
		if err := e.service.UpdateCustomer(&customer); err != nil {
			log.Printf("Error updating PUT /customer/record...(%v)", err.Error())
			http.Error(w, "Error updating PUT /customer/record", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if err := e.service.DeleteCustomer(id); err != nil {
			log.Printf("Error deleting DELETE /customer/record...(%v)", err.Error())
			http.Error(w, "Error deleting DELETE /customer/record", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
