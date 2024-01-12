package cli

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5/config"
	"github.com/koki-develop/gh-q/internal/git"
)

type Client struct {
	root      string
	gitClient *git.Client
}

func NewClient() (*Client, error) {
	// load gitconfig
	gitcfg, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		return nil, err
	}

	// get root
	root, err := getRoot(gitcfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		root:      root,
		gitClient: git.NewClient(),
	}, nil
}

func getRoot(gitcfg *config.Config) (string, error) {
	// priority: env > gitconfig > default

	// env
	if r := os.Getenv("GHQ_ROOT"); r != "" {
		return r, nil
	}

	// gitconfig
	if r := gitcfg.Raw.Section("ghq").Option("root"); r != "" {
		return r, nil
	}

	// default
	h, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(h, "ghq"), nil
}

func (c *Client) path(owner, repo string) string {
	return filepath.Join(c.root, "github.com", owner, repo)
}
