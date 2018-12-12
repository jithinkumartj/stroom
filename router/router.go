package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stroom/auth"

	"github.com/gorilla/mux"
)

var authenticator auth.AuthenticateClient

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, status := authenticator.Authenticate(r)
		if session == nil {
			http.Error(w, "Forbidden", status)
			return
		}
		log.Printf("Authenticated user")
		next.ServeHTTP(w, r)
	})
}

func Init() {

	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("GET")
	router.Handle("/closed_endpoint", AuthenticationMiddleware(http.HandlerFunc(closed_endpoint)))
	router.HandleFunc("/open_endpoint", open_endpoint)
	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func closed_endpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I guess you have logged in!")
}
func open_endpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is public!")
}
func login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authenticator.Login(w, r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged in")
}

func logout(w http.ResponseWriter, r *http.Request) {
	authenticator.Logout(w, r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Logged out")
}
