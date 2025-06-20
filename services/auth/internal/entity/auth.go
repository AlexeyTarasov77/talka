package entity

type OAuthProviderId int

const (
	GOOGLE_OAUTH_PROVIDER OAuthProviderId = iota
	GITHUB_OAUTH_PROVIDER
)

func (p OAuthProviderId) String() string {
	return []string{"Google", "GitHub"}[p]
}

type OAuthProviderInfo struct {
	ID   OAuthProviderId
	Name string
}

var OAuthSupportedProviders = []OAuthProviderId{
	GOOGLE_OAUTH_PROVIDER, GITHUB_OAUTH_PROVIDER,
}
