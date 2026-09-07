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

	mux.Handle("/product/category", utils.AuthMiddleware(http.HandlerFunc(handler.CallCategory)))
	mux.Handle("/product/category/list", utils.AuthMiddleware(http.HandlerFunc(handler.CallCategoryList)))
	mux.Handle("/product/category/new", utils.AuthMiddleware(http.HandlerFunc(handler.CallCategoryNew)))
	mux.Handle("/product/category/record", utils.AuthMiddleware(http.HandlerFunc(handler.CallCategoryRecord)))
	mux.Handle("/proveedor", utils.AuthMiddleware(http.HandlerFunc(handler.CallProveedor)))
	mux.Handle("/proveedor/slice", utils.AuthMiddleware(http.HandlerFunc(handler.CallProveedorSlice)))
	mux.Handle("/proveedor/new", utils.AuthMiddleware(http.HandlerFunc(handler.CallProveedorNew)))
	mux.Handle("/proveedor/record", utils.AuthMiddleware(http.HandlerFunc(handler.CallProveedorRecord)))
	mux.Handle("/product", utils.AuthMiddleware(http.HandlerFunc(handler.CallProduct)))
	mux.Handle("/product/slice", utils.AuthMiddleware(http.HandlerFunc(handler.CallProductSlice)))
	mux.Handle("/product/new", utils.AuthMiddleware(http.HandlerFunc(handler.CallProductNew)))
	mux.Handle("/product/record", utils.AuthMiddleware(http.HandlerFunc(handler.CallProductRecord)))
}
