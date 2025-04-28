package gogithubapp

import (
	"context"

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
func (c *Client) CreateFromTemplate(ctx context.Context, tmplOwner, tmplRepo, newName string) (*githubapi.Repository, error) {
	tmpl := &githubapi.TemplateRepoRequest{Name: githubapi.String(newName)}
	repo, _, err := c.client.Repositories.CreateFromTemplate(ctx, tmplOwner, tmplRepo, tmpl)
	if err != nil {
		return nil, err
	}
	return repo, nil
}
