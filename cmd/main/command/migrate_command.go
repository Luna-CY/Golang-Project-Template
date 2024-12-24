package command

import "github.com/spf13/cobra"

func NewMigrateCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "migrate",
		Args: cobra.NoArgs,
	}

	command.AddCommand(NewMigrateUpCommand(), NewMigrateDownCommand())

	return command
}
