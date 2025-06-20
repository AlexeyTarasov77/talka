package oauth

import (
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways"
	"github.com/AlexeyTarasov77/messanger.users/internal/gateways/oauth/providers"
)

type Providers struct {
	Google gateways.OAuthProvider
}

func New(clientID, clientSecret, redirectURL string) *Providers {
	return &Providers{
		Google: oauth_providers.NewGoogleProvider(clientID, clientSecret, redirectURL),
	}
}
