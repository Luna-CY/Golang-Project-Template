package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code      Code   `json:"code" validate:"required"`           // code, 0 if success, non-zero otherwise
	RequestId string `json:"request_id" validate:"required"`     // request id, used for tracking
	Msg       string `json:"msg" validate:"required"`            // message, OK on success, error message otherwise
	Data      any    `json:"data,omitempty" validate:"optional"` // data
}

type BaseDataList[T any] struct {
	Page  int   `json:"page" validate:"required"`  // page number
	Size  int   `json:"size" validate:"required"`  // number of items per page
	Total int64 `json:"total" validate:"required"` // total count
	Data  []T   `json:"data" validate:"required"`  // data list
}

func Success(c *gin.Context, data any) {
	var response Response
	response.Code = Ok
	response.RequestId = c.Request.Header.Get("X-Request-ID")
	response.Msg = "OK"
	response.Data = data

	c.JSON(http.StatusOK, response)
}

func Failure(c *gin.Context, code Code, message string) {
	var response Response
	response.Code = code
	response.RequestId = c.Request.Header.Get("X-Request-ID")
	response.Msg = message

	c.JSON(http.StatusOK, response)
}

func FailureWithData(c *gin.Context, code Code, message string, data any) {
	var response Response
	response.Code = code
	response.RequestId = c.Request.Header.Get("X-Request-ID")
	response.Msg = message
	response.Data = data

	c.JSON(http.StatusOK, response)
}

type Redirect struct {
	To string
}

func RedirectTo(to string) Redirect {
	return Redirect{To: to}
}
