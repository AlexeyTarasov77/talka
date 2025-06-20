package oauth

import "github.com/AlexeyTarasov77/messanger.users/internal/gateways/oauth/providers"

type Providers struct {
	Google oauth_providers.Interface
}

func New(clientID, clientSecret, redirectURL string) *Providers {
	return &Providers{
		Google: oauth_providers.NewGoogleProvider(clientID, clientSecret, redirectURL),
	}
}
