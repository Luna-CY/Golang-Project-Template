package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/i18n"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"github.com/Luna-CY/Golang-Project-Template/server/http/request"
	"github.com/Luna-CY/Golang-Project-Template/server/http/response"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Field1 string `json:"field1" validate:"required" maximum:"255" binding:"required,max=255"`                         // field1
	Field2 uint64 `json:"field2" validate:"required" minimum:"1" maximum:"999999" binding:"required,min=1,max=999999"` // field2
	Field3 bool   `json:"field3" validate:"required" binding:"omitempty"`                                              // field3
	Field4 int    `json:"field4" validate:"required" enums:"1,2,3" binding:"required,oneof=1 2 3"`                     // field4, enums: 1 - A, 2 - B, 3 - C
}

// Create
// @Tags category/example
// @Summary create
// @Param param body CreateRequest true "request body"
// @Success 200 {object} response.Response{} "successful. click to expand response structure"
// @Router /example/create [post]
func (cls *Example) Create(c *gin.Context) (response.Code, any, errors.I18nError) {
	var body = CreateRequest{}
	if err := request.ShouldBindJSON(c, &body); nil != err {
		return response.InvalidParams, nil, errors.NewI18n(i18n.CommonIdInvalidRequest, err.Relation(errors.ErrorInvalidRequest("SHGWHE_LE.E_LE.C_TE.292902")))
	}

	var ctx = contextutil.NewContextWithGin(c)
	if _, err := cls.example.CreateExample(ctx, body.Field1, body.Field2, body.Field3, model.ExampleEnumFieldType(body.Field4)); nil != err {
		return response.ServerInternalError, nil, errors.NewI18n(i18n.CommonIdServerInternalError, err.Relation(errors.ErrorServerInternalError("SHGWHE_LE.E_LE.C_TE.342933")))
	}

	return response.Ok, nil, nil
}
