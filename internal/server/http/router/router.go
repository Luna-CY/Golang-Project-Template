package router

import (
	"errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext/icontextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/ierror"
	"github.com/Luna-CY/Golang-Project-Template/internal/server/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Wrapper(handler func(*gin.Context) (response.Code, any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code, res, err = handler(c)
		if response.Ok != code || nil != err {
			var message = err.Error()

			var ie *ierror.Error
			if errors.As(err, &ie) {
				message = ie.I18n(icontextutil.NewContextWithGin(c))
			}

			if nil != res {
				response.FailureWithData(c, code, message, res)

				return
			}

			response.Failure(c, code, message)

			return
		}

		if redirect, ok := res.(response.Redirect); ok {
			c.Redirect(http.StatusFound, redirect.To)

			return
		}

		response.Success(c, res)
	}
}
