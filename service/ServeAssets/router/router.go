package router

import (
	"net/http"
)

const GlobalAssetsPathWeb string = "/assets/"

// Serve Global Assets
func ServeGlobalAssetsStaticFiles(mux *http.ServeMux) {

	fs := http.FileServer(http.Dir("./assets/"))
	mux.Handle(GlobalAssetsPathWeb, http.StripPrefix(GlobalAssetsPathWeb, fs))
}
