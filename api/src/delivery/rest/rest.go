package rest

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/lib/twitter"
)

var sessionKey string

func init() {
	sessionKey = config.Conf.SessionKey
}

var store = sessions.NewCookieStore([]byte(sessionKey))
var authSessionName = "as"

func authHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start authHandler")
	config := twitter.GetConnect()

	rt, err := config.RequestTemporaryCredentials(nil, "http://localhost:8080/twitter/callback", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("rt: %+v\n", rt)

	session, err := store.Get(r, authSessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["request_token"] = rt.Token
	session.Values["request_token_secret"] = rt.Secret

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := config.AuthorizationURL(rt, nil)
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

var notExistErrFormat = "%s doese not exist in the session"

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start callbackHandler")

	session, err := store.Get(r, authSessionName)
	if err != nil {
		log.Printf("failed to get session, the cause was %v", err)
		handleSessionError(w, err)
		return
	}

	requestToken, ok := session.Values["request_token"].(string)
	if !ok {
		log.Printf(notExistErrFormat, "request_token")
		handleNotMatchToken(w)
		return
	}

	requestTokenSecret, ok := session.Values["request_token_secret"].(string)
	if !ok {
		log.Printf(notExistErrFormat, "request_token_secret")
		handleNotMatchToken(w)
		return
	}

	if requestToken != r.URL.Query().Get("oauth_token") {
		log.Println("request oauth_token not equal request_token_secret in session", "request_token_secret")
		handleNotMatchToken(w)
		return
	}

	at, err := twitter.GetAccessToken(
		&oauth.Credentials{
			Token:  requestToken,
			Secret: requestTokenSecret,
		},
		r.URL.Query().Get("oauth_verifier"),
	)
	if err != nil {
		panic(err)
	}

	account := twitter.Account{}
	if err = twitter.GetMe(at, &account); err != nil {
		panic(err)
	}

	fmt.Println(strings.Repeat("#", 20))
	fmt.Printf("\n\naccount: %#v\n\n", account)
	fmt.Println(strings.Repeat("#", 20))

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Success!!")
}

func handleSessionError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "server error")
}

func handleNotMatchToken(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "token did not match")
}

// NewRoutes will initialize the all resources endpoint
func NewRoutes(r *chi.Mux) {
	r.Get("/twitter/oauth", authHandler)
	r.Get("/twitter/callback", callbackHandler)
}
