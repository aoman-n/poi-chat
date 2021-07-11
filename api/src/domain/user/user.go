package user

type User struct {
	ID        int
	UID       string
	Name      string
	AvatarURL string
	Provider  Provider
}

type Provider string

const (
	ProviderTwitter Provider = "twitter"
	ProviderGuest   Provider = "guest"
)
