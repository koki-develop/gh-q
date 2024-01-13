package cli

import (
	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/cli/go-gh/v2/pkg/config"
	"github.com/koki-develop/gh-q/internal/git"
)

func (c *Client) auth() (*git.Auth, error) {
	h, _ := auth.DefaultHost()
	tkn, _ := auth.TokenForHost(h)
	if tkn == "" {
		return nil, nil
	}

	username, err := c.GetUsername()
	if err != nil {
		return nil, err
	}

	return &git.Auth{
		Username: username,
		Token:    tkn,
	}, nil
}

func (c *Client) GetUsername() (string, error) {
	ghcfg, err := config.Read()
	if err != nil {
		return "", err
	}

	u, err := ghcfg.Get([]string{"hosts", "github.com", "user"})
	if err != nil {
		return "", err
	}

	return u, nil
}
