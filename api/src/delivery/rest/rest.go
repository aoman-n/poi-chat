package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/laster18/poi/api/src/config"
)

var (
	errGetCredentials     = errors.New("cannot get credentials")
	errInvalidSession     = errors.New("invalid session")
	errSaveSession        = errors.New("cannot save session")
	errInvalidCredentials = errors.New("invalid credentials")
	errImageSave          = errors.New("failed to save image")
)

// NewRoutes will initialize the all resources endpoint
func NewRoutes(r *chi.Mux) {
	r.Get("/twitter/oauth", twitterOauthHandler)
	r.Get("/twitter/callback", twitterCallbackHandler)
	r.Post("/guest-login", guestLoginHandler)
	r.Get("/logout", logoutHandler)
}

type ResponseErr struct {
	Message string `json:"message"`
}

func getResponseErrJSON(msg string) (string, error) {
	resErr := ResponseErr{
		Message: msg,
	}

	bs, err := json.Marshal(resErr)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func handleInternalServerStrErr(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "internal server error")
}

func handleInvalidSessionErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, fmt.Sprintf("%s: %s", errInvalidSession.Error(), err.Error()))
}

func handleSaveOrRemoveSessionErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, fmt.Sprintf("%s: %s", errSaveSession.Error(), err.Error()))
}

func handleNotMatchTokenErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, fmt.Sprintf("%s: %s", errInvalidCredentials.Error(), err.Error()))
}

func handleImageSaveErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, fmt.Sprintf("%s: %s", errImageSave.Error(), err.Error()))
}

func handleValidationErr(w http.ResponseWriter, err error) {
	resJSON, err := getResponseErrJSON(err.Error())
	if err != nil {
		handleInternalServerStrErr(w)
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, resJSON)
}

func handleRedirectRoot(w http.ResponseWriter) {
	w.Header().Set("location", config.Conf.FrontBaseURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
