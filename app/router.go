package app
import (
	"code.google.com/p/gorilla/mux"
	"net/http"
)

var(
	router *mux.Router
)

func routerSetup(){
	router = mux.NewRouter()
	router.Handle("/", indexHandler)
	router.Handle("/tweet/{id}", showHandler)
	router.Handle("/search", indexHandler)

	http.Handle("/", router)
	http.Handle("/public/", http.FileServer(http.Dir(rootPath)))
}
