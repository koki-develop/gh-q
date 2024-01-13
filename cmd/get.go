package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/koki-develop/gh-q/internal/cli"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use: fmt.Sprintf("get %s/%s|%s",
		color.New(color.Italic).Sprint("OWNER"), color.New(color.Italic).Sprint("REPO"),
		color.New(color.Italic).Sprint("REPO"),
	),
	Aliases: []string{"g"},
	Short:   "Clone repository",
	Long:    "Clone repository.",
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
			owner, err = c.GetUsername()
			if err != nil {
				return err
			}
			repo = segs[0]
		}
		if len(segs) == 2 {
			owner, repo = segs[0], segs[1]
		}

		if err := c.Get(owner, repo); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
