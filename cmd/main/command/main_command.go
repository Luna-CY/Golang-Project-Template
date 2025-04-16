package command

import (
	"os"

	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration/loader"
	"github.com/Luna-CY/Golang-Project-Template/internal/i18n"
	"github.com/Luna-CY/Golang-Project-Template/internal/runtime"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
)

func NewMainCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "main",
		Args: cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := loader.LoadConfig("config", &configuration.Configuration); nil != err {
				cmd.PrintErrf("error loading config: %v\n", err)

				os.Exit(1)
			}

			if err := i18n.Init(); nil != err {
				cmd.PrintErrf("Error init i18n: %v\n", err)
			}

			if configuration.Configuration.Sentry.Enable && "" != configuration.Configuration.Sentry.Dsn {
				hostname, _ := os.Hostname()

				option := sentry.ClientOptions{
					Release:          runtime.Release,
					AttachStacktrace: true,
					Debug:            configuration.Configuration.Debug,
					ServerName:       hostname,
					Environment:      runtime.GetEnvironment(),
					Dsn:              configuration.Configuration.Sentry.Dsn,
				}

				if err := sentry.Init(option); err != nil {
					cmd.PrintErrf("Error init sentry: %v\n", err)

					os.Exit(1)
				}
			}

			// Perform any necessary setup or initialization here
		},
	}

	// Add subcommands here
	command.AddCommand(NewMigrateCommand(), NewServerCommand(), NewGenerateCommand())

	return command
}
