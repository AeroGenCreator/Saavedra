package router

import (
	"Saavedra/service/Login/api"
	"Saavedra/service/Login/service"
	"Saavedra/service/Login/store"
	"database/sql"
	"net/http"
)

func Assambler(mux *http.ServeMux, db *sql.DB) {

	store := store.New(db)
	service := service.New(store)
	handler := api.New(service)

	mux.HandleFunc("/login", handler.CallLogin)
	mux.HandleFunc("/", handler.CallRoot)
}
