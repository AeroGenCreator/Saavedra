package router

import (
	"Saavedra/service/Login/api"
	"Saavedra/service/Login/service"
	"Saavedra/service/Login/store"
	utils "Saavedra/utils"
	"database/sql"
	"net/http"
)

func Assambler(mux *http.ServeMux, db *sql.DB) {

	store := store.New(db)
	service := service.New(store)
	handler := api.New(service)

	mux.Handle("/login", utils.LowLevelMiddleware(http.HandlerFunc(handler.CallLogin)))
	mux.Handle("/refresh", utils.LowLevelMiddleware(http.HandlerFunc(handler.RefreshToken)))
	mux.HandleFunc("/", handler.CallRoot)
}
