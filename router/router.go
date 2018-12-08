package router

import(
	"fmt"
    "html"
    "net/http"
    "github.com/gorilla/mux"
)

func route(w http.ResponseWriter, r *http.Request) {
	
}

func Init() {
	router := mux.NewRouter();
	router.HandleFunc("/", Index)
}