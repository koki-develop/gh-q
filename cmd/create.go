package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/koki-develop/gh-q/internal/cli"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: fmt.Sprintf("create %s/%s|%s...",
		color.New(color.Italic).Sprint("OWNER"), color.New(color.Italic).Sprint("REPO"),
		color.New(color.Italic).Sprint("REPO"),
	),
	Aliases: []string{"c", "new", "n"},
	Short:   "Create a new repository on local",
	Long:    "Create a new repository on local.",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := cli.NewClient()
		if err != nil {
			return err
		}

		for _, arg := range args {
			owner, repo, err := cli.ParseOwnerRepo(arg)
			if err != nil {
				return err
			}
			if err := c.Create(owner, repo); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
