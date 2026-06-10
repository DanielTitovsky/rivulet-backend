package domain

type ProviderType string

const (
	ProviderGoogle ProviderType = "google"
	ProviderYandex ProviderType = "yandex"
	ProviderGit    ProviderType = "git"
)

type OAuthUser struct {
	Provider          ProviderType
	ProviderUserId    string
	ProviderUserEmail string
	EmailVerified     bool
	Name              string
	GivenName         string
	FamilyName        string
}
