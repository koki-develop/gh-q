package cmd

import (
	"errors"
	"strings"

	"github.com/koki-develop/gh-q/internal/cli"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:     "clone",
	Aliases: []string{"c"},
	Short:   "Clone repository",
	Long:    "Clone repository.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		segs := strings.Split(args[0], "/")
		if len(segs) != 2 {
			return errors.New("invalid repository name")
		}
		owner, repo := segs[0], segs[1]

		c, err := cli.NewClient()
		if err != nil {
			return err
		}

		if err := c.Clone(owner, repo); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
