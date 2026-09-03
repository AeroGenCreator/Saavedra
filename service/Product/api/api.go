package api

import (
	"Saavedra/service/Product/service"
	"Saavedra/service/Product/types"
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

// === MENU ===

// ROUTE: /product/menu -> Renders HTML for SERVICE MENU.
func ProductMenu(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tpl, err := template.ParseFiles("service/Product/views/productMenu.html")
		if err != nil {
			log.Printf("Error parsing template '/product/menu'...(%v)", err.Error())
			http.Error(w, "Error parsing template '/product/menu'", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, nil)
		if err != nil {
			log.Printf("Error rendering template '/product/menu'...(%v)", err.Error())
			http.Error(w, "Error rendering template '/product/menu'", http.StatusInternalServerError)
			return
		}
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}

// === MATERIAL ===

// ROUTE: /product/material -> Renders HTML for LIST VIEW.
func (e *EndpointHandler) CallMaterial(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tpl, err := template.ParseFiles("service/Product/views/productMaterial.html")
		if err != nil {
			log.Printf("Error parsing HTML /product/material...(%v)", err.Error())
			http.Error(w, "Error parsing HTML /product/material", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, nil)
		if err != nil {
			log.Printf("Error rendering HTML /product/material...(%v)", err.Error())
			http.Error(w, "Error rendering HTML /product/material", http.StatusInternalServerError)
			return
		}
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		log.Print("Invalid Method /product/material")
		http.Error(w, "Invalid Method /product/material", http.StatusMethodNotAllowed)
		return
	}
}

// ROUTE: /product/material/list -> RETURNS rows
func (e *EndpointHandler) CallMaterialList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		page := r.URL.Query().Get("page")
		if page == "" {
			page = "1"
		}
		intPage, err := strconv.Atoi(page)
		if err != nil {
			intPage = 1
		}
		materialSlice, err := e.service.ListMaterial(intPage)
		if err != nil {
			log.Printf("Error fetching /product/material/list...(%v)", err.Error())
			http.Error(w, "Error fetching /product/material/list", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(materialSlice)
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		log.Print("Invalid Method /product/material")
		http.Error(w, "Invalid Method /product/material", http.StatusMethodNotAllowed)
		return
	}
}

// ROUTE: /material/new
func (e *EndpointHandler) CallMaterialNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tpl, err := template.ParseFiles("service/Product/views/productMaterialNew.html")
		if err != nil {
			log.Printf("Error parsing /product/material/new...(%v)", err.Error())
			http.Error(w, "Error parsing /product/material/new", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, nil)
		if err != nil {
			log.Printf("Error rendering /product/material/new...(%v)", err.Error())
			http.Error(w, "Error rendering /product/material/new", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var material types.Material
		if err := json.NewDecoder(r.Body).Decode(&material); err != nil {
			log.Printf("Error decoding r.Body '/product/material/new'...(%v)", err.Error())
			http.Error(w, "Error decoding r.Body '/product/material/new'", http.StatusInternalServerError)
			return
		}
		_, err := e.service.CreateMaterial(&material)
		if err != nil {
			log.Printf("Error creating '/product/material/new'...(%v)", err.Error())
			http.Error(w, "Error creating '/product/material/new'", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		log.Print("Invalid Method /product/material/new")
		http.Error(w, "Invalid Method /product/material/new", http.StatusMethodNotAllowed)
		return
	}
}

// ROUTE: /product/material/record
func (e *EndpointHandler) CallMaterialRecord(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		record, err := e.service.ReadMaterial(id)
		if err == types.ErrNoRecord {
			log.Printf("Error query /product/material/record...(%v)", err.Error())
			http.Error(w, "Error query /product/material/record", http.StatusForbidden)
			return
		} else if err != nil {
			log.Printf("Error query /product/material/record...(%v)", err.Error())
			http.Error(w, "Error query /product/material/record", http.StatusInternalServerError)
			return
		}
		tpl, err := template.ParseFiles("service/Product/views/productMaterialRecord.html")
		if err != nil {
			log.Printf("Error parsing /product/material/record...(%v)", err.Error())
			http.Error(w, "Error parsing /product/material/record", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, record)
		if err != nil {
			log.Printf("Error rendering /product/material/record...(%v)", err.Error())
			http.Error(w, "Error rendering /product/material/record", http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		var material types.MaterialStr
		if err := json.NewDecoder(r.Body).Decode(&material); err != nil {
			log.Printf("Error parsing /product/material/record...(%v)", err.Error())
			http.Error(w, "Error parsing /product/material/record", http.StatusInternalServerError)
			return
		}
		record, err := e.service.UpdateMaterial(&material)
		if err != nil {
			log.Printf("Error query /product/material/record...(%v)", err.Error())
			http.Error(w, "Error query /product/material/record", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(record)
	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		if err := e.service.DeleteMaterial(id); err != nil {
			log.Printf("Error query /product/material/record...(%v)", err.Error())
			http.Error(w, "Error query /product/material/record", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodHead:
		w.WriteHeader(http.StatusOK)
	default:
		log.Print("Invalid Method /product/material/record")
		http.Error(w, "Invalid Method /product/material/record", http.StatusMethodNotAllowed)
		return
	}
}

// === PROVEEDOR ===

// === PRODUCT ===
