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
	cfg, err := config.LoadConfig(config.GlobalScope)
	if err != nil {
		return nil, err
	}

	// load root directory
	ghq := cfg.Raw.Section("ghq")
	root := ghq.Option("root")
	if root == "" {
		// if not set, use `$HOME/ghq`
		r, err := defaultRoot()
		if err != nil {
			return nil, err
		}
		root = r
	}

	return &Client{
		root:      root,
		gitClient: git.NewClient(),
	}, nil
}

func defaultRoot() (string, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(h, "ghq"), nil
}
