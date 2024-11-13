package command

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration/loader"
	"github.com/Luna-CY/Golang-Project-Template/internal/i18n"
	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "server",
		Args: cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := loader.LoadConfig("config", &configuration.Configuration); nil != err {
				cmd.PrintErrf("Error loading config: %v\n", err)

				return
			}

			if err := i18n.Init(); nil != err {
				cmd.PrintErrf("Error init i18n: %v\n", err)

				return
			}

			// Perform any necessary setup or initialization here
		},
	}

	command.AddCommand(NewHttpCommand())

	return command
}
