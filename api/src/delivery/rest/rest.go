package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

var (
	errGetCredentials     = errors.New("cannot get credentials")
	errInvalidSession     = errors.New("invalid session")
	errSaveSession        = errors.New("cannot save session")
	errInvalidCredentials = errors.New("invalid credentials")
)

// NewRoutes will initialize the all resources endpoint
func NewRoutes(r *chi.Mux) {
	r.Get("/twitter/oauth", twitterOauthHandler)
	r.Get("/twitter/callback", twitterCallbackHandler)
	r.Get("/logout", logoutHandler)
}

func handleInvalidSessionErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, fmt.Sprintf("%s: %s", errInvalidSession.Error(), err.Error()))
}

func handleSaveOrRemoveSessionErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, fmt.Sprintf("%s: %s", errSaveSession.Error(), err.Error()))
}

func handleNotMatchToken(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, fmt.Sprintf("%s: %s", errInvalidCredentials.Error(), err.Error()))
}
