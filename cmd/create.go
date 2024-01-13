package cmd

import (
	"errors"
	"strings"

	"github.com/koki-develop/gh-q/internal/cli"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create OWNER/REPO",
	Aliases: []string{"c"},
	Short:   "Create a new repository",
	Long:    "Create a new repository.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cli.NewClient()
		if err != nil {
			return err
		}

		var owner string
		var repo string

		segs := strings.Split(args[0], "/")
		if len(segs) > 2 {
			return errors.New("invalid repository name")
		}
		if len(segs) == 1 {
			repo = segs[0]
			owner, err = c.GetUsername()
			if err != nil {
				return err
			}
		}
		if len(segs) == 2 {
			owner, repo = segs[0], segs[1]
		}

		if err := c.Create(owner, repo); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
