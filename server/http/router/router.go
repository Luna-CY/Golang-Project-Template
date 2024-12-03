package router

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	response2 "github.com/Luna-CY/Golang-Project-Template/server/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Wrapper(handler func(*gin.Context) (response2.Code, any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code, res, err = handler(c)
		if response2.Ok != code || nil != err {
			var message = err.Error()

			var ie *errors.Error
			if errors.As(err, &ie) {
				message = ie.I18n(contextutil.NewContextWithGin(c))
			}

			if nil != res {
				response2.FailureWithData(c, code, message, res)

				return
			}

			response2.Failure(c, code, message)

			return
		}

		if redirect, ok := res.(response2.Redirect); ok {
			c.Redirect(http.StatusFound, redirect.To)

			return
		}

		response2.Success(c, res)
	}
}
