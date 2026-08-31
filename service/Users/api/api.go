package api

import (
	"Saavedra/service/Users/service"
	"net/http"
)

type EndpointHandler struct {
	service service.Service
}

func New(s service.Service) *EndpointHandler {
	return &EndpointHandler{service: s}
}

func (e EndpointHandler) CallUsers(w http.ResponseWriter, r *http.Request) {

}
