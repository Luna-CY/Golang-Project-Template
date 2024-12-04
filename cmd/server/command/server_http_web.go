package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	middleware2 "github.com/Luna-CY/Golang-Project-Template/server/http/middleware"
	"github.com/Luna-CY/Golang-Project-Template/server/http/web"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

func NewHttpWebCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:  "web",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			gin.SetMode(gin.ReleaseMode)
			var engine = gin.New()

			// register common middleware
			engine.Use(middleware2.CustomGinRecovery(), middleware2.RequestId())

			if err := engine.SetTrustedProxies(configuration.Configuration.Server.Http.Web.GinTrustedProxies); nil != err {
				cmd.PrintErrf("error setting trusted proxies: %v\n", err)

				os.Exit(1)
			}

			if configuration.Configuration.Debug {
				gin.SetMode(gin.DebugMode)
				engine.Use(gin.Logger())
			}

			// register routes and middleware here
			web.Register(engine)

			// register under maintenance middleware if enabled in configuration.Configuration.Server.Http.Web.UnderMaintenance
			if configuration.Configuration.Server.Http.Web.UnderMaintenance {
				engine.Use(middleware2.UnderMaintenance())
			}

			// set listening address if not set in configuration.Configuration.Server.Http.Web.Listen, default to ":8000"
			if "" == configuration.Configuration.Server.Http.Web.Listen {
				configuration.Configuration.Server.Http.Web.Listen = ":8000"
			}

			// manual create and start server
			var server = http.Server{
				Addr:    configuration.Configuration.Server.Http.Web.Listen,
				Handler: engine,
			}

			go func() {
				fmt.Printf("starting server on %s\n", configuration.Configuration.Server.Http.Web.Listen)

				if err := server.ListenAndServe(); nil != err {
					if !errors.Is(err, http.ErrServerClosed) {
						cmd.PrintErrf("error starting server: %v\n", err)

						os.Exit(1)
					}
				}
			}()

			// wait context cancellation signal
			<-cmd.Context().Done()

			// stop server gracefully
			var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); nil != err {
				cmd.PrintErrf("error shutting down server: %v\n", err)
			}

			fmt.Println("server stopped")
		},
	}

	return command
}
