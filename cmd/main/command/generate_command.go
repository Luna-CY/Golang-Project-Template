package command

import "github.com/spf13/cobra"

func NewGenerateCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "generate",
		Args: cobra.NoArgs,
	}

	command.AddCommand(NewGenerateDaoCommand())

	return command
}
