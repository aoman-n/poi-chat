package session

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/domain/user"
)

var sessionSecret string

func init() {
	sessionSecret = config.Conf.SessionKey
}

var store = sessions.NewCookieStore([]byte(sessionSecret))

const (
	authSessionName = "auth"
	userSessionName = "user"
)

type AuthSession struct {
	sess *sessions.Session
}

func GetAuthSession(r *http.Request) (*AuthSession, error) {
	sess, err := store.Get(r, authSessionName)
	if err != nil {
		return nil, err
	}

	return &AuthSession{sess: sess}, nil
}

func (a *AuthSession) SetCredentials(token, secret string) {
	a.sess.Values["request_token"] = token
	a.sess.Values["request_token_secret"] = secret
}

func (a *AuthSession) Save(r *http.Request, w http.ResponseWriter) error {
	return a.sess.Save(r, w)
}

func (a *AuthSession) GetCredentials() (token string, secret string, err error) {
	token, ok := a.sess.Values["request_token"].(string)
	if !ok {
		return "", "", fmt.Errorf("not found token in session")
	}

	secret, ok = a.sess.Values["request_token_secret"].(string)
	if !ok {
		return "", "", fmt.Errorf("not found secret in session")
	}

	return token, secret, nil
}

func (a *AuthSession) RemoveCredentials(r *http.Request, w http.ResponseWriter) error {
	a.sess.Options.MaxAge = -1
	return a.sess.Save(r, w)
}

type UserSession struct {
	sess *sessions.Session
}

const (
	idKey     = "user_id"
	uidKey    = "user_uid"
	nameKey   = "user_name"
	avatarKey = "avatar_url"
)

func GetUserSession(r *http.Request) (*UserSession, error) {
	sess, err := store.Get(r, userSessionName)
	if err != nil {
		return nil, err
	}

	return &UserSession{sess: sess}, nil
}

func (s *UserSession) SetUserUID(u *user.User) {
	s.sess.Values[uidKey] = u.UID
}

func (s *UserSession) Save(r *http.Request, w http.ResponseWriter) error {
	// TODO: add domain option
	s.sess.Options.HttpOnly = true
	s.sess.Options.SameSite = http.SameSiteStrictMode
	// s.sess.Options.SameSite = http.SameSiteLaxMode
	s.sess.Options.Path = "/"
	if config.IsProd() {
		s.sess.Options.Secure = true
	}
	return s.sess.Save(r, w)
}

func (s *UserSession) IsNew() bool {
	return s.sess.IsNew
}

func (s *UserSession) GetUserUID() (string, error) {
	uid, ok := s.sess.Values[uidKey].(string)
	if !ok {
		return "", fmt.Errorf("not found user uid in session")
	}

	return uid, nil
}

func (s *UserSession) RemoveUser(r *http.Request, w http.ResponseWriter) error {
	s.sess.Options.MaxAge = -1
	return s.sess.Save(r, w)
}
