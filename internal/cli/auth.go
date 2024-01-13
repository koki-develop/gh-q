package cli

import (
	"github.com/cli/go-gh/v2/pkg/auth"
	"github.com/koki-develop/gh-q/internal/git"
)

func (c *Client) auth() (*git.Auth, error) {
	h, _ := auth.DefaultHost()
	tkn, _ := auth.TokenForHost(h)
	if tkn == "" {
		return nil, nil
	}

	username, err := GetUsername()
	if err != nil {
		return nil, err
	}

	return &git.Auth{
		Username: username,
		Token:    tkn,
	}, nil
}
