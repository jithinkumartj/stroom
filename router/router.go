package router

import(
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
)

type authenticator interface {
	Authenticate(string) error
}

var auth authenticator 

func route(w http.ResponseWriter, r *http.Request) {
	session := r.Header.Get("x-session-id")
	err := auth.Authenticate(session)
	if err != nil {
		fmt.Println(err)
	}
}

func Init(authClient authenticator) {
	
	router := mux.NewRouter();
	router.HandleFunc("/", route)
	auth = authClient
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}