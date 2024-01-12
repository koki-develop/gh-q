package git

import (
	"github.com/cli/go-gh/v2/pkg/config"
)

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
