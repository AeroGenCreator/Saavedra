package router

import (
	"net/http"
	"path/filepath"
	"runtime"
)

const GlobalAssetsPathWeb string = "/assets/"

func ServeGlobalAssetsStaticFiles(mux *http.ServeMux) {

	// Caller file 'main.go'
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information available")
	}

	// Extracts the directory from the file path
	rootDir := filepath.Dir(filename)

	// Endpoint to serve global assets
	fs := http.FileServer(http.Dir(rootDir + "assets"))
	mux.Handle(GlobalAssetsPathWeb, http.StripPrefix(GlobalAssetsPathWeb, fs))
}
