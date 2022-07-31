package cmd

import (
	"github.com/nathancoleman/render-sdk-go"
	"github.com/spf13/cobra"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Interact with Render services",
}

var servicesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your Render services",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := render.NewClient(render.DefaultConfig())
		if err != nil {
			return err
		}

		services, err := client.Services().List(cmd.Context())
		if err != nil {
			return err
		}

		for _, service := range services {
			cmd.Println(service.ID, "\t", service.Name)
		}

		return nil
	},
}

func init() {
	servicesCmd.AddCommand(servicesListCmd)
}
