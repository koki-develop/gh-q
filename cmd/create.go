package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/koki-develop/gh-q/internal/cli"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: fmt.Sprintf("create %s/%s|%s",
		color.New(color.Italic).Sprint("OWNER"), color.New(color.Italic).Sprint("REPO"),
		color.New(color.Italic).Sprint("REPO"),
	),
	Aliases: []string{"c"},
	Short:   "Create a new repository on local",
	Long:    "Create a new repository on local.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner, repo, err := cli.ParseOwnerRepo(args[0])
		if err != nil {
			return err
		}

		c, err := cli.NewClient()
		if err != nil {
			return err
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
