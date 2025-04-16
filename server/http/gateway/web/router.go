package web

import (
	"io"
	"net/http"
	"time"

	"github.com/Luna-CY/Golang-Project-Template/internal/build"
	"github.com/Luna-CY/Golang-Project-Template/internal/docs"
	"github.com/Luna-CY/Golang-Project-Template/internal/runtime"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	var root = engine.Group("/api")

	// enable CORS for development environment, disable it in production environment.
	if !runtime.IsDevelopment() {
		var cc = cors.Config{
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: false, MaxAge: 12 * time.Hour, AllowAllOrigins: true,
		}

		root.Use(cors.New(cc))
	}

	// register docs api if build.Docs is true
	_ = build.Docs && doc(root)

	{
		// example
		// var example = handler.NewExample()
		// var group = root.Group("example")

		// group.POST("create", router.Wrapper(example.Create))
		// group.POST("update", router.Wrapper(example.Update))
		// group.POST("list", router.Wrapper(example.List))
		// group.POST("detail", router.Wrapper(example.Detail))
	}
}

func doc(group *gin.RouterGroup) bool {
	// re-doc.html
	group.GET("/doc", func(ctx *gin.Context) {
		html, err := docs.ReDoc.Open("re-doc.html")
		if nil != err {
			ctx.AbortWithStatus(500)
			return
		}

		if _, err := io.Copy(ctx.Writer, html); nil != err {
			ctx.AbortWithStatus(500)
			return
		}

		ctx.Status(http.StatusOK)
	})

	group.GET("/main_swagger.json", func(ctx *gin.Context) {
		html, err := docs.Swagger.Open("main_swagger.json")
		if nil != err {
			ctx.AbortWithStatus(500)
			return
		}

		if _, err := io.Copy(ctx.Writer, html); nil != err {
			ctx.AbortWithStatus(500)
			return
		}

		ctx.Status(http.StatusOK)
	})

	return true
}
