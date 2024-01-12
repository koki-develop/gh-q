package git

import "github.com/cli/go-gh/v2/pkg/api"

func (c *Client) Viewer() (string, error) {
	gql, err := api.DefaultGraphQLClient()
	if err != nil {
		return "", err
	}

	var q struct {
		Viewer struct {
			Login string
		}
	}
	if err := gql.Query("viewer", &q, nil); err != nil {
		return "", err
	}

	return q.Viewer.Login, nil
}
