package twitter

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"os"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/laster18/poi/api/src/config"
)

var (
	tempCredKey  string
	tokenCredKey string
)

type Account struct {
	ID              string `json:"id_str"`
	ScreenName      string `json:"screen_name"` // acountId
	ProfileImageURL string `json:"profile_image_url_https"`
	Name            string `json:"name"` // displayName
}

func init() {
	tempCredKey = config.Conf.Twitter.APIKey
	tokenCredKey = config.Conf.Twitter.SecretKey
}

func GetConnect() *oauth.Client {
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
		TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
		Credentials: oauth.Credentials{
			Token:  tempCredKey,
			Secret: tokenCredKey,
		},
	}
}

func GetAccessToken(rt *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
	oc := GetConnect()
	at, _, err := oc.RequestToken(nil, rt, oauthVerifier)

	return at, err
}

func GetMe(at *oauth.Credentials, user *Account) error {
	oc := GetConnect()

	v := url.Values{}
	v.Set("include_email", "true")

	resp, err := oc.Get(nil, at, "https://api.twitter.com/1.1/account/verify_credentials.json", v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return errors.New("Twitter is unavailable")
	}

	if resp.StatusCode >= 400 {
		return errors.New("Twitter request is invalid")
	}

	// debug print
	var r io.Reader = resp.Body
	r = io.TeeReader(r, os.Stdout)

	err = json.NewDecoder(r).Decode(user)
	if err != nil {
		return err
	}

	return nil

}
