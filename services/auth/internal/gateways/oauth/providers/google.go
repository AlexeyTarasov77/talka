package oauth_providers

import (
	"net/http"

	"github.com/AlexeyTarasov77/messanger.users/internal/entity"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Google struct {
	cfg    *oauth2.Config
	client *http.Client
}

func NewGoogleProvider(clientID, clientSecret, redirectURL string) *Google {
	oauthConfig := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return &Google{cfg: oauthConfig}
}

func (g *Google) GetAuthURL(stateToken string) string {
	return g.cfg.AuthCodeURL(stateToken)
}
func (g *Google) GetAccessToken(ctx context.Context, authCode string) (string, error) {
	token, err := g.cfg.Exchange(ctx, authCode)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

func (g *Google) FetchUserData(ctx context.Context, accessToken string) (*entity.User, error) {
	// client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
	return nil, nil
}
