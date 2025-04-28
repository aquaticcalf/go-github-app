package gogithubapp

import (
	"context"
	"fmt"

	githubapi "github.com/google/go-github/v50/github"
)

// ListRepositories lists repositories accessible to the client
func (c *Client) ListRepositories(ctx context.Context) ([]*githubapi.Repository, error) {
	var all []*githubapi.Repository
	opt := &githubapi.ListOptions{PerPage: 50}
	for {
		repos, resp, err := c.client.Repositories.List(ctx, "", &githubapi.RepositoryListOptions{ListOptions: *opt})
		if err != nil {
			return nil, err
		}
		all = append(all, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return all, nil
}

// CreateRepository creates a new repository under the authenticated actor
func (c *Client) CreateRepository(ctx context.Context, name, description string, private bool) (*githubapi.Repository, error) {
	ren := &githubapi.Repository{Name: githubapi.String(name), Description: githubapi.String(description), Private: githubapi.Bool(private)}
	repo, _, err := c.client.Repositories.Create(ctx, "", ren)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

// CreateFromTemplate makes a new repository from a template
func (c *Client) CreateFromTemplate(ctx context.Context, tmplOwner, tmplRepo, newName, owner string) (*githubapi.Repository, error) {

	tmpl := &githubapi.TemplateRepoRequest{
		Owner: githubapi.String(owner),
		Name:  githubapi.String(newName),
		// TODO : Description: githubapi.String(""),
		// TODO : Private:     githubapi.Bool(false),
	}

	repo, _, err := c.client.Repositories.CreateFromTemplate(ctx, tmplOwner, tmplRepo, tmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to create repository '%s' owned by '%s' from template %s/%s: %w", newName, owner, tmplOwner, tmplRepo, err)
	}
	return repo, nil
}

// GetInstallation retrieves details about a specific installation using the client's authentication.
func (c *Client) GetInstallation(ctx context.Context, installationID int64) (*githubapi.Installation, error) {
	installation, _, err := c.client.Apps.GetInstallation(ctx, installationID)
	if err != nil {
		return nil, fmt.Errorf("failed to get installation details for ID %d: %w", installationID, err)
	}
	return installation, nil
}
