package cmd

import (
	"os"

	"github.com/nathancoleman/render-sdk-go"
	"github.com/spf13/cobra"
)

func init() {
	deploysCmd.AddCommand(deploysListCmd)
}

var deploysCmd = &cobra.Command{
	Use:   "deploys",
	Short: "Interact with Render deploys",
}

var deploysListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your Render deployments",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := render.NewClient(render.DefaultConfig())
		if err != nil {
			return err
		}

		deploys, err := client.Deploys(args[0]).List(cmd.Context())
		if err != nil {
			return err
		}

		templates.ExecuteTemplate(os.Stdout, "deploys.tmpl", struct{ Deploys []render.Deploy }{deploys})
		return nil
	},
}
