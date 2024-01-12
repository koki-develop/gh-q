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

	viewer, err := c.gitClient.GetUsername()
	if err != nil {
		return nil, err
	}

	return &git.Auth{
		Username: viewer,
		Token:    tkn,
	}, nil
}
