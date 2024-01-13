package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/koki-develop/gh-q/internal/cli"
	"github.com/spf13/cobra"
)

var (
	flagRemoveForce bool
)

var removeCmd = &cobra.Command{
	Use: fmt.Sprintf("remove %s/%s|%s...",
		color.New(color.Italic).Sprint("OWNER"), color.New(color.Italic).Sprint("REPO"),
		color.New(color.Italic).Sprint("REPO"),
	),
	Aliases: []string{"rm"},
	Short:   "Remove repository from local",
	Long:    "Remove repository from local.",
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

			if err := c.Remove(owner, repo, cli.WithRemoveForce(flagRemoveForce)); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().BoolVarP(&flagRemoveForce, "force", "f", false, "Remove without confirmation")
}
