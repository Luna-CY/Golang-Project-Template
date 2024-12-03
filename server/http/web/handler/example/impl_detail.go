package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/server/http/request"
	"github.com/Luna-CY/Golang-Project-Template/server/http/response"
	"github.com/gin-gonic/gin"
)

type DetailRequest struct {
	Id uint64 `json:"id" validate:"required" binding:"required"` // record id
}

type DetailResponse struct {
	Id         uint64 `json:"id" validate:"required"`                   // record id
	Field1     string `json:"field1" validate:"required"`               // field1
	Field2     uint64 `json:"field2" validate:"required"`               // field2
	Field3     bool   `json:"field3" validate:"required"`               // field3
	Field4     int    `json:"field4" validate:"required" enums:"1,2,3"` // field4, enums: 1 - A, 2 - B, 3 - C
	CreateTime int64  `json:"create_time" validate:"required"`          // create time
	UpdateTime int64  `json:"update_time" validate:"required"`          // update time
}

// Detail
// @Tags category/example
// @Summary detail
// @Param param body DetailRequest true "request body"
// @Success 200 {object} response.Response{data=DetailResponse{}} "successful. click to expand response structure"
// @Router /example/detail [post]
func (cls *Example) Detail(c *gin.Context) (response.Code, any, error) {
	var body = DetailRequest{}
	if err := request.ShouldBindJSON(c, &body); nil != err {
		return response.InvalidParams, nil, err
	}

	var ctx = contextutil.NewContextWithGin(c)
	data, err := cls.example.GetExampleById(ctx, body.Id, false)
	if nil != err {
		if errors.Is(err, errors.ErrorRecordNotFound) {
			return response.InvalidParams, nil, errors.New("example record not found: %d", body.Id)
		}

		return response.ServerInternalError, nil, err
	}

	var res = pointer.Default(DetailResponse{
		Id:         data.Id,
		Field1:     *data.Field1,
		Field2:     *data.Field2,
		Field3:     *data.Field3,
		Field4:     int(*data.Field4),
		CreateTime: *data.CreateTime,
		UpdateTime: *data.UpdateTime,
	})

	return response.Ok, res, nil
}
