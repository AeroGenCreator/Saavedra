package api

import (
	"html/template"
	"log"
	"net/http"
)

// TRANSPORT "PRODUCT MENU"
// ROUTE: /product/menu
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
