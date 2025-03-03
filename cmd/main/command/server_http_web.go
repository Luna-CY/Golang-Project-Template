package command

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"github.com/Luna-CY/Golang-Project-Template/internal/http"
	"github.com/spf13/cobra"
)

func NewHttpWebCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "web",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// set listening address if not set in configuration.Configuration.Server.Http.Web.Listen, default to ":8000"
			if "" == configuration.Configuration.Server.Http.Web.Listen {
				configuration.Configuration.Server.Http.Web.Listen = ":8080"
			}

			http.Listen(cmd, configuration.Configuration.Server.Http.Web.Listen, configuration.Configuration.Server.Http.Web.GinTrustedProxies, configuration.Configuration.Server.Http.Web.UnderMaintenance)
		},
	}

	return command
}
