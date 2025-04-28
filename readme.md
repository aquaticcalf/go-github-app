# gogithubapp

`gogithubapp` is a go client library that unifies github oauth2 user flows and github app installation workflows, providing advanced abstractions for common operations.

---

## Installation

```bash
go get github.com/aquaticcalf/gogithubapp
```

## Usage

### 1. oauth2 ( user flow )

```go
import (
  "context"
  "golang.org/x/oauth2"
  "github.com/aquaticcalf/gogithubapp"
)

// 1. build config
cfg := gogithubapp.OAuthConfig(
  os.Getenv("GITHUB_CLIENT_ID"),
  os.Getenv("GITHUB_CLIENT_SECRET"),
  "https://yourapp/callback",
  []string{"repo", "workflow"},
)

// 2. redirect user to
state := "random-state"
authURL := cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)

// 3. on callback
tok, err := gogithubapp.ExchangeToken(context.Background(), cfg, code)
client := gogithubapp.NewClientFromToken(context.Background(), tok)

// 4. use client
repos, err := client.ListRepositories(context.Background())
```

### 2. github app ( installation flow )

```go
import (
  "context"
  "github.com/aquaticcalf/gogithubapp"
)

// provide raw or base64-encoded PEM key
appID := int64(123456)
installID := int64(789012)
pemKey := []byte(os.Getenv("GITHUB_PRIVATE_KEY"))

client, err := gogithubapp.NewClientFromApp(
  context.Background(), appID, installID, pemKey,
)
```

## api reference

### repositories

```go
client.ListRepositories(ctx)              // []*github.Repository
client.CreateRepository(ctx, name, desc, private)
client.CreateFromTemplate(ctx, owner, tmplRepo, newName)
```

### branches

```go
client.CreateBranch(ctx, owner, repo, newBranch, baseBranch)
client.ListBranches(ctx, owner, repo)
```

### commits

```go
client.CommitMultipleFiles(ctx, owner, repo, branch, message, files) // files map[string][]byte
client.ListCommits(ctx context.Context, owner, repo, opts)
```

### issues & pr

```go
client.CreateIssue(ctx, owner, repo, title, body)
client.ListIssues(ctx, owner, repo, opts)
client.CreatePullRequest(ctx, owner, repo, title, head, base, body)
client.ListPullRequests(ctx, owner, repo, opts)
client.MergePullRequest(ctx, owner, repo, prNumber, msg)
client.CreateComment(ctx, owner, repo, number, comment)
client.AddLabels(ctx, owner, repo, number, labels)
```

---

## license

[MIT](license.md)
