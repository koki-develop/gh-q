package cli

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	ghconfig "github.com/cli/go-gh/v2/pkg/config"
	"github.com/go-git/go-git/v5/config"
	"github.com/koki-develop/gh-q/internal/git"
)

type Client struct {
	root      string
	gitClient *git.Client
}

func NewClient() (*Client, error) {
	// get root
	root, err := GetRoot()
	if err != nil {
		return nil, err
	}

	return &Client{
		root:      root,
		gitClient: git.NewClient(),
	}, nil
}

func (c *Client) path(owner, repo string) string {
	return filepath.Join(c.root, "github.com", owner, repo)
}

func LoadGitConfig() (*config.Config, error) {
	return config.LoadConfig(config.GlobalScope)
}

func LoadGhConfig() (*ghconfig.Config, error) {
	return ghconfig.Read()
}

func GetRoot() (string, error) {
	// priority: env > gitconfig > default

	// env
	if r := os.Getenv("GHQ_ROOT"); r != "" {
		return r, nil
	}

	// gitconfig
	gitcfg, err := LoadGitConfig()
	if err != nil {
		return "", err
	}
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

func GetUsername() (string, error) {
	// priority: env > gitconfig > gh config

	// env
	if u := os.Getenv("GHQ_USER"); u != "" {
		return u, nil
	}

	// gitconfig
	gitcfg, err := LoadGitConfig()
	if err != nil {
		return "", err
	}
	if u := gitcfg.Raw.Section("ghq").Option("user"); u != "" {
		return u, nil
	}

	// gh config
	ghcfg, err := LoadGhConfig()
	if err != nil {
		return "", err
	}
	u, err := ghcfg.Get([]string{"hosts", "github.com", "user"})
	if err != nil {
		return "", err
	}
	return u, nil
}

func ParseOwnerRepo(s string) (string, string, error) {
	s = strings.TrimPrefix(s, "github.com/")

	segs := strings.Split(s, "/")
	if len(segs) == 1 {
		repo := segs[0]
		owner, err := GetUsername()
		if err != nil {
			return "", "", err
		}
		return owner, repo, nil
	}
	if len(segs) == 2 {
		return segs[0], segs[1], nil
	}

	return "", "", errors.New("invalid repository name")
}
