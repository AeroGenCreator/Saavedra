package router

import (
	"Saavedra/service/Quote/api"
	"Saavedra/utils"
	"database/sql"
	"net/http"
)

func Assambler(mux *http.ServeMux, db *sql.DB) {
	mux.Handle("/quote/menu", utils.AuthMiddleware(http.HandlerFunc(api.CallQuoteMenu)))
	mux.Handle("/quote/new", utils.AuthMiddleware(http.HandlerFunc(api.CallQuoteNew)))
}
