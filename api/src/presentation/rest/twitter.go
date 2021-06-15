package rest

import (
	"log"
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/util/session"
	"github.com/laster18/poi/api/src/util/twitter"
)

func twitterOauthHandler(w http.ResponseWriter, r *http.Request) {
	twitterClient := twitter.GetConnect()

	rt, err := twitterClient.RequestTemporaryCredentials(nil, config.Conf.Twitter.CallbackURI, nil)
	if err != nil {
		log.Print("failed to request credentials err:", err)
		panic(err)
	}

	session, err := session.GetAuthSession(r)
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

func twitterCallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start callbackHandler")

	sess, err := session.GetAuthSession(r)
	if err != nil {
		log.Printf("failed to get auth session, the cause was %v", err)
		handleInvalidSessionErr(w, err)
		return
	}

	token, secret, err := sess.GetCredentials()
	if err != nil {
		log.Print(err)
		handleNotMatchTokenErr(w, err)
		return
	}

	if token != r.URL.Query().Get("oauth_token") {
		log.Println("request oauth_token not equal request_token_secret in session", "request_token_secret")
		handleNotMatchTokenErr(w, err)
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

	userSession, err := session.GetUserSession(r)
	if err != nil {
		log.Printf("failed to get user session, the cause was %v", err)
		handleInvalidSessionErr(w, err)
		return
	}

	userSession.SetUser(&session.User{
		ID:        account.ID,
		Name:      account.Name,
		AvatarURL: account.ProfileImageURL,
	})
	if err := userSession.Save(r, w); err != nil {
		log.Print("failed to set user to session err:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := sess.RemoveCredentials(r, w); err != nil {
		handleSaveOrRemoveSessionErr(w, err)
		return
	}

	w.Header().Set("location", config.Conf.FrontBaseURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
