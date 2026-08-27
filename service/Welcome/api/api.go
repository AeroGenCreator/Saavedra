package api

import (
	"html/template"
	"net/http"
)

func TransportWelcomeTemplate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("Backend/Services/Welcome/views/welcome.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
}
