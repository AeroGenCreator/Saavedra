package router

import (
	"Saavedra/service/Welcome/api"
	utils "Saavedra/utils"
	"net/http"
)

func Assambler(mux *http.ServeMux) {

	// AuthMiddleware Protects backend from requests
	mux.Handle("/welcome", utils.AuthMiddleware(http.HandlerFunc(api.TransportWelcomeTemplate)))
}
