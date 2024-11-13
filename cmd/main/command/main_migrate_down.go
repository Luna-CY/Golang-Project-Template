package command

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/spf13/cobra"
)

func NewMigrateDownCommand() *cobra.Command {
	var path string
	var number int
	var force bool

	var command = &cobra.Command{
		Use:   "down",
		Short: "Apply database migrations down to the specified version or the previous version",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			var db, err = sql.Open("mysql", configuration.Configuration.Database.Mysql.Dsn)
			if nil != err {
				cmd.PrintErrf("Error connecting to the database: %v\n", err)

				return
			}

			defer func() {
				if err := db.Close(); nil != err {
					cmd.PrintErrf("Error closing the database connection: %v\n", err)
				}
			}()

			driver, err := mysql.WithInstance(db, &mysql.Config{})
			if nil != err {
				cmd.PrintErrf("Error creating the database driver: %v\n", err)

				return
			}

			defer func() {
				if err := driver.Close(); nil != err {
					cmd.PrintErrf("Error closing the database driver: %v\n", err)
				}
			}()

			mi, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", path), "mysql", driver)
			if nil != err {
				cmd.PrintErrf("Error creating the migration instance: %v\n", err)

				return
			}

			if number > 0 {
				if err := mi.Steps(-number); nil != err && !errors.Is(err, migrate.ErrNoChange) {
					cmd.PrintErrf("Error revert migrations: %v\n", err)

					return
				}
			} else {
				if !force {
					fmt.Println("Danger: This will revert all migrations. place use --force flag to proceed.")

					return
				}

				if err := mi.Down(); nil != err && !errors.Is(err, migrate.ErrNoChange) {
					cmd.PrintErrf("Error revert migrations: %v\n", err)

					return
				}
			}

			version, dirty, err := mi.Version()
			if nil != err {
				cmd.PrintErrf("Error getting migration version: %v\n", err)

				return
			}

			// Print success message
			cmd.Printf("Database migrations revert successfully. current version: %d, dirty: %t\n", version, dirty)
		},
	}

	command.Flags().StringVarP(&path, "path", "p", "./migration", "Path to the migration directory. default:./migration")
	command.Flags().IntVarP(&number, "number", "n", 1, "Apply migrations down to the specified version. default: apply the previous version")
	command.Flags().BoolVarP(&force, "force", "f", false, "Force apply all migration. default: false")

	return command
}
