package auth

import (
	"net/http"
	"stroom/auth/session_store"

	"github.com/gorilla/sessions"
)

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type AuthenticateClient struct{}

func (AuthenticateClient) Authenticate(r *http.Request) (*sessions.Session, int) {
	store := session_store.Store
	session, _ := store.Get(r, "login-session")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return nil, http.StatusUnauthorized
	}
	return session, http.StatusOK
}

func (AuthenticateClient) Login(w http.ResponseWriter, r *http.Request) {
	store := session_store.Store
	session, _ := store.Get(r, "login-session")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func (AuthenticateClient) Logout(w http.ResponseWriter, r *http.Request) {
	store := session_store.Store
	session, _ := store.Get(r, "login-session")
	session.Options.MaxAge = -1
	session.Save(r, w)
}
