package command

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration/loader"
	"github.com/spf13/cobra"
)

func NewMainCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "main",
		Args: cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := loader.LoadConfig("config", &configuration.Configuration); nil != err {
				cmd.PrintErrf("error loading config: %v\n", err)

				return
			}

			// Perform any necessary setup or initialization here
		},
	}

	// Add subcommands here
	command.AddCommand(NewMigrateCommand())

	return command
}
