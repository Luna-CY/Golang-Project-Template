package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/i18n"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao/option"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"github.com/Luna-CY/Golang-Project-Template/server/http/request"
	"github.com/Luna-CY/Golang-Project-Template/server/http/response"
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
func (cls *Example) List(c *gin.Context) (response.Code, any, errors.I18nError) {
	var body = ListRequest{BaseListRequest: request.BaseListRequest{Page: 1, Size: 20}}
	if err := request.ShouldBindJSON(c, &body); nil != err {
		return response.InvalidParams, nil, errors.NewI18n(i18n.CommonIdInvalidRequest, err.Relation(errors.ErrorInvalidRequest("SHGWHE_LE.E_LE.L_ST.373016")))
	}

	var options = []option.ExampleOption{option.ExampleOptionWithOrderDefault()}
	if 0 != body.Field4 {
		options = append(options, option.ExampleOptionWithField4(model.ExampleEnumFieldType(body.Field4)))
	}

	// create context with gin context
	// not allow use gin context in internal service
	var ctx = contextutil.NewContextWithGin(c)

	// search data
	total, data, err := cls.example.ListBySimpleCondition(ctx, body.Page, body.Size, options...)
	if nil != err {
		return response.ServerInternalError, nil, errors.NewI18n(i18n.CommonIdServerInternalError, err.Relation(errors.ErrorServerInternalError("SHGWHE_LE.E_LE.L_ST.523027")))
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
