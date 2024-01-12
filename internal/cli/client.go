package cli

import "github.com/koki-develop/gh-q/internal/git"

type Client struct {
	gitClient *git.Client
}

func NewClient() *Client {
	return &Client{
		gitClient: git.NewClient(),
	}
}
