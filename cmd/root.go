package cmd

import (
	"os"
	"text/template"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var templates = template.Must(
	template.New("").
		Funcs(template.FuncMap{
			"date":  func(t time.Time) string { return t.Format(time.RFC1123) },
			"green": color.GreenString,
		}).
		ParseGlob("tmpl/*.tmpl"))

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "render",
	Short: "Interact with Render",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(
		servicesCmd,
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.render-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
