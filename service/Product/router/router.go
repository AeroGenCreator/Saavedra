package router

import (
	"Saavedra/service/Product/api"
	"Saavedra/utils"
	"net/http"
)

func Assambler(mux *http.ServeMux) {

	// AuthMiddleware Protects backend from requests
	mux.Handle("/product/menu", utils.AuthMiddleware(http.HandlerFunc(api.ProductMenu)))
}
