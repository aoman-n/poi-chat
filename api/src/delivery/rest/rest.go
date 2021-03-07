package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/go-chi/chi"
	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/delivery"
	"github.com/laster18/poi/api/src/lib/twitter"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	twitterClient := twitter.GetConnect()

	rt, err := twitterClient.RequestTemporaryCredentials(nil, config.Conf.Twitter.CallbackURI, nil)
	if err != nil {
		log.Print("failed to request credentials err:", err)
		panic(err)
	}

	session, err := delivery.GetAuthSession(r)
	if err != nil {
		log.Print("failed to get auth session err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.SetCredentials(rt.Token, rt.Secret)
	if err := session.Save(r, w); err != nil {
		log.Print("failed to set credentials to session err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := twitterClient.AuthorizationURL(rt, nil)
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

var notExistErrFormat = "%s doese not exist in the session"

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start callbackHandler")

	session, err := delivery.GetAuthSession(r)
	if err != nil {
		log.Printf("failed to get auth session, the cause was %v", err)
		handleSessionError(w, err)
		return
	}

	token, secret, err := session.GetCredentials()
	if err != nil {
		log.Print(err)
		handleNotMatchToken(w)
		return
	}

	if token != r.URL.Query().Get("oauth_token") {
		log.Println("request oauth_token not equal request_token_secret in session", "request_token_secret")
		handleNotMatchToken(w)
		return
	}

	at, err := twitter.GetAccessToken(
		&oauth.Credentials{
			Token:  token,
			Secret: secret,
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

	userSession, err := delivery.GetUserSession(r)
	if err != nil {
		log.Printf("failed to get user session, the cause was %v", err)
		handleSessionError(w, err)
		return
	}

	userSession.SetUser(&delivery.User{
		ID:        account.ID,
		Name:      account.Name,
		AvatarURL: account.ProfileImageURL,
	})
	if err := userSession.Save(r, w); err != nil {
		log.Print("failed to set user to session err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
