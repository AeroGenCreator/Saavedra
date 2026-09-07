package router

import (
	"Saavedra/service/Customer/api"
	"Saavedra/service/Customer/service"
	"Saavedra/service/Customer/store"
	"Saavedra/utils"
	"database/sql"
	"net/http"
)

func Assambler(mux *http.ServeMux, db *sql.DB) {
	store := store.New(db)
	service := service.New(store)
	handler := api.New(service)

	mux.Handle("/customer", utils.AuthMiddleware(http.HandlerFunc(handler.CallCustomer)))
	mux.Handle("/customer/slice", utils.AuthMiddleware(http.HandlerFunc(handler.CallCustomerList)))
	mux.Handle("/customer/new", utils.AuthMiddleware(http.HandlerFunc(handler.CallCustomerNew)))
	mux.Handle("/customer/record", utils.AuthMiddleware(http.HandlerFunc(handler.CallCustomerRecord)))
}
