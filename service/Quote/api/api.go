package api

import (
	"html/template"
	"log"
	"net/http"
)

type EndpointHandler struct {
}

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
