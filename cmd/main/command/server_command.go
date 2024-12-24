package command

import (
	"github.com/spf13/cobra"
)

func NewServerCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "server",
		Args: cobra.NoArgs,
	}

	command.AddCommand(NewHttpCommand())

	return command
}
