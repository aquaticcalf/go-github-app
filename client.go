package gogithubapp

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	githubapp "github.com/bradleyfalzon/ghinstallation/v2"
	githubapi "github.com/google/go-github/v50/github"
	oauth2 "golang.org/x/oauth2"
)

// Client wraps a go-github client for unified api operations
type Client struct {
	client *githubapi.Client
}

// NewClientFromToken creates a client using an oauth2 token
func NewClientFromToken(ctx context.Context, token *oauth2.Token) *Client {
	ts := oauth2.StaticTokenSource(token)
	httpClient := oauth2.NewClient(ctx, ts)
	return &Client{client: githubapi.NewClient(httpClient)}
}

// NewClientFromApp creates a client using a gitHub app installation
// privateKeyPEM may be raw or base64-encoded PEM
func NewClientFromApp(ctx context.Context, appID, installID int64, privateKeyPEM []byte) (*Client, error) {
	// decode if base64
	if len(privateKeyPEM) > 0 && privateKeyPEM[0] != '-' {
		decoded, err := base64.StdEncoding.DecodeString(string(privateKeyPEM))
		if err != nil {
			return nil, fmt.Errorf("invalid base64 key: %w", err)
		}
		privateKeyPEM = decoded
	}
	itr, err := githubapp.New(http.DefaultTransport, appID, installID, privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to init installation transport: %w", err)
	}
	return &Client{client: githubapi.NewClient(&http.Client{Transport: itr})}, nil
}
