package http

import (
	"context"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/server/http/gateway/web"
	"github.com/Luna-CY/Golang-Project-Template/server/http/middleware"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

// Listen start http listener
func Listen(cmd *cobra.Command, listen string, proxies []string, underMaintenance bool) {
	gin.SetMode(gin.ReleaseMode)
	var engine = gin.New()

	// register common middleware
	engine.Use(middleware.CustomGinRecovery(), middleware.RequestId())

	if err := engine.SetTrustedProxies(proxies); nil != err {
		cmd.PrintErrf("error setting trusted proxies: %v\n", err)

		os.Exit(1)
	}

	if configuration.Configuration.Debug {
		gin.SetMode(gin.DebugMode)
		engine.Use(gin.Logger())
	}

	// register ping route
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// register routes and middleware here
	web.Register(engine)

	// register under maintenance middleware if enabled on configuration.Configuration.Server.Http.Web.UnderMaintenance
	if underMaintenance {
		engine.Use(middleware.UnderMaintenance())
	}

	// register sentry middleware if it is enabled on configuration.Configuration.Sentry.Enable
	if configuration.Configuration.Sentry.Enable {
		engine.Use(sentrygin.New(sentrygin.Options{Repanic: true, WaitForDelivery: true, Timeout: 10 * time.Second}))
	}

	// manual create and start server
	var server = http.Server{
		Addr:    listen,
		Handler: engine,
	}

	go func() {
		fmt.Printf("starting server on %s\n", listen)

		if err := server.ListenAndServe(); nil != err {
			if !errors.Is(err, http.ErrServerClosed) {
				cmd.PrintErrf("error starting server: %v\n", err)

				os.Exit(1)
			}
		}
	}()

	// wait context cancellation signal
	<-cmd.Context().Done()

	// flush sentry buffer
	sentry.Flush(5 * time.Second)

	// stop server gracefully
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); nil != err {
		cmd.PrintErrf("error shutting down server: %v\n", err)
	}

	fmt.Println("server stopped")
}
