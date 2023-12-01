package cmd

import (
	"github.com/smartfood-capstone/backend/internal/app"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cli := &cobra.Command{}
	cli.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Starting server",
		Long:  `Starting server`,
		Run: func(c *cobra.Command, _ []string) {
			cmd := app.NewStartCmd()
			cmd.Start()
		},
	})

	return cli
}
