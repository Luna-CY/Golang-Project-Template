package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext/icontextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/server/http/request"
	"github.com/Luna-CY/Golang-Project-Template/internal/server/http/response"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	request.BaseListRequest

	Field4 int `json:"field4" validate:"optional" enums:"0,1,2,3" binding:"omitempty,oneof=0 1 2 3"` // filter by field4, enums: 0 - All, 1 - A, 2 - B, 3 - C
}

type ListItem struct {
	Field1 string `json:"field1" validate:"required"`               // field1
	Field2 uint64 `json:"field2" validate:"required"`               // field2
	Field3 bool   `json:"field3" validate:"required"`               // field3
	Field4 int    `json:"field4" validate:"required" enums:"1,2,3"` // field4, enums: 1 - A, 2 - B, 3 - C
}

// List
// @Tags category/example
// @Summary list for data
// @Param param body ListRequest true "request body"
// @Success 200 {object} response.Response{data=response.BaseDataList[ListItem]{}} "successful. click to expand response structure"
// @Router /example/list [post]
func (cls *Example) List(c *gin.Context) (response.Code, any, error) {
	var body = ListRequest{BaseListRequest: request.BaseListRequest{Page: 1, Size: 20}}
	if err := request.ShouldBindJSON(c, &body); nil != err {
		return response.InvalidParams, nil, err
	}

	var field4 *model.ExampleEnumFieldType
	if 0 != body.Field4 {
		field4 = pointer.New(model.ExampleEnumFieldType(body.Field4))
	}

	// create context with gin context
	// not allow use gin context in internal service
	var ctx = icontextutil.NewContextWithGin(c)

	// search data
	total, data, err := cls.example.ListBySimpleCondition(ctx, field4, body.Page, body.Size)
	if nil != err {
		return response.ServerInternalError, nil, err
	}

	// response data
	var res = pointer.Default(response.BaseDataList[ListItem]{Total: total, Page: body.Page, Size: body.Size})
	for _, example := range data {
		var item = ListItem{
			Field1: *example.Field1,
			Field2: *example.Field2,
			Field3: *example.Field3,
			Field4: int(*example.Field4),
		}

		res.Data = append(res.Data, item)
	}

	return response.Ok, res, nil
}
