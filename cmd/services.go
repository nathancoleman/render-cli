package cmd

import (
	"os"

	"github.com/nathancoleman/render-sdk-go"
	"github.com/spf13/cobra"
)

func init() {
	servicesCmd.AddCommand(servicesListCmd, servicesGetCmd)
}

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

		templates.ExecuteTemplate(os.Stdout, "services.tmpl", struct{ Services []render.Service }{services})

		return nil
	},
}

var servicesGetCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"retrieve"},
	Short:   "Get information about a particular Render service",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := render.NewClient(render.DefaultConfig())
		if err != nil {
			return err
		}

		service, err := client.Services().Get(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		templates.ExecuteTemplate(os.Stdout, "service.tmpl", service)
		return nil
	},
}
