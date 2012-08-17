package app

import (
	"code.google.com/p/gorilla/mux"
	"net/http"
)

var (
	router *mux.Router
)

func routerSetup() {
	router = mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/tweet/{id}", showHandler)
	router.HandleFunc("/search", indexHandler)

	http.Handle("/", router)
	http.Handle("/public/", http.FileServer(http.Dir(rootPath)))
}
