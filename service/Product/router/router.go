package router

import (
	"Saavedra/service/Product/api"
	"Saavedra/service/Product/service"
	"Saavedra/service/Product/store"
	"Saavedra/utils"
	"database/sql"
	"net/http"
)

func Assambler(mux *http.ServeMux, db *sql.DB) {

	store := store.New(db)
	service := service.New(store)
	handler := api.New(service)

	// AuthMiddleware Protects backend from requests
	mux.Handle("/product/menu", utils.AuthMiddleware(http.HandlerFunc(api.ProductMenu)))

	mux.Handle("/product/material", utils.AuthMiddleware(http.HandlerFunc(handler.CallMaterial)))
	mux.Handle("/product/material/list", utils.AuthMiddleware(http.HandlerFunc(handler.CallMaterialList)))
	mux.Handle("/product/material/new", utils.AuthMiddleware(http.HandlerFunc(handler.CallMaterialNew)))
	mux.Handle("/product/material/record", utils.AuthMiddleware(http.HandlerFunc(handler.CallMaterialRecord)))
	mux.Handle("/proveedor", utils.AuthMiddleware(http.HandlerFunc(handler.CallProveedor)))
	mux.Handle("/proveedor/slice", utils.AuthMiddleware(http.HandlerFunc(handler.CallProveedorSlice)))
}
