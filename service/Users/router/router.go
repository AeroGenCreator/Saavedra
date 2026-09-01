package router

import (
	"Saavedra/service/Users/api"
	"Saavedra/service/Users/service"
	"Saavedra/service/Users/store"
	"Saavedra/utils"
	"database/sql"
	"net/http"
)

func Assambler(mux *http.ServeMux, db *sql.DB) {

	store := store.New(db)
	service := service.New(store)
	handler := api.New(service)

	mux.Handle("/users", utils.AuthMiddleware(http.HandlerFunc(handler.CallUsers)))
	mux.Handle("/users/records", utils.AuthMiddleware(http.HandlerFunc(handler.CallUsersRecords)))

}
