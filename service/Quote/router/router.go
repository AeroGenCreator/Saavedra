package router

import (
	"Saavedra/service/Quote/api"
	"Saavedra/service/Quote/service"
	"Saavedra/service/Quote/store"
	"Saavedra/utils"
	"database/sql"
	"net/http"
)

func Assambler(mux *http.ServeMux, db *sql.DB) {

	storing := store.New(db)
	servicing := service.New(storing)
	handler := api.New(servicing)

	mux.Handle("/quote/menu", utils.AuthMiddleware(http.HandlerFunc(api.CallQuoteMenu)))
	mux.Handle("/quote/new", utils.AuthMiddleware(http.HandlerFunc(api.CallQuoteNew)))
	mux.Handle("/quote/many2one", utils.AuthMiddleware(http.HandlerFunc(handler.CallMany2One)))
}
