package gogithubapp

import (
	"context"

	oauth2 "golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
)

// OAuthConfig constructs an oauth2 configuration for user auth via github
func OAuthConfig(clientID, clientSecret, redirectURL string, scopes []string) *oauth2.Config {
	if len(scopes) == 0 {
		scopes = []string{"repo", "workflow"}
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     oauth2github.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
	}
}

// ExchangeToken exchanges the oauth2 code for a token
func ExchangeToken(ctx context.Context, cfg *oauth2.Config, code string) (*oauth2.Token, error) {
	return cfg.Exchange(ctx, code)
}
