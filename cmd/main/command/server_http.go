package command

import "github.com/spf13/cobra"

func NewHttpCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "http",
		Args: cobra.NoArgs,
	}

	command.AddCommand(NewHttpWebCommand())

	return command
}
