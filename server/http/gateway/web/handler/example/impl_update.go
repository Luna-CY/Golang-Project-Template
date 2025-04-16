package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/i18n"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"github.com/Luna-CY/Golang-Project-Template/server/http/request"
	"github.com/Luna-CY/Golang-Project-Template/server/http/response"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Id     uint64 `json:"id" validate:"required" binding:"required"`                                                   // record id
	Field1 string `json:"field1" validate:"required" maximum:"255" binding:"required,max=255"`                         // field1
	Field2 uint64 `json:"field2" validate:"required" minimum:"1" maximum:"999999" binding:"required,min=1,max=999999"` // field2
	Field3 bool   `json:"field3" validate:"required" binding:"omitempty"`                                              // field3
	Field4 int    `json:"field4" validate:"required" enums:"1,2,3" binding:"required,oneof=1 2 3"`                     // field4, enums: 1 - A, 2 - B, 3 - C
}

// Update
// @Tags category/example
// @Summary update
// @Param param body UpdateRequest true "request body"
// @Success 200 {object} response.Response{} "successful. click to expand response structure"
// @Router /example/update [post]
func (cls *Example) Update(c *gin.Context) (response.Code, any, errors.I18nError) {
	var body = UpdateRequest{}
	if err := request.ShouldBindJSON(c, &body); nil != err {
		return response.InvalidParams, nil, errors.NewI18n(i18n.CommonIdInvalidRequest, err.Relation(errors.ErrorInvalidRequest("SHGWHE_LE.E_LE.U_TE.313038")))
	}

	var ctx = contextutil.NewContextWithGin(c)
	example, err := cls.example.GetExampleById(ctx, body.Id)
	if nil != err {
		if err.IsType(errors.ErrorTypeRecordNotFound) {
			return response.InvalidParams, nil, errors.NewI18n(i18n.CommonIdRecordNotFound, err.Relation(errors.New(errors.ErrorTypeInvalidRequest, "SHGWHE_LE.E_LE.U_TE.383049", "example record not found: %d", body.Id)))
		}

		return response.ServerInternalError, nil, errors.NewI18n(i18n.CommonIdServerInternalError, err.Relation(errors.ErrorServerInternalError("SHGWHE_LE.E_LE.U_TE.413052")))
	}

	// if need transaction
	//if err := cls.example.WithTransaction(ctx, func(ctx context.Context) errors.Error {
	//	if err := cls.example.UpdateExample(ctx, example, pointer.New(body.Field1), pointer.New(body.Field2), pointer.New(body.Field3), pointer.New(model.ExampleEnumFieldType(body.Field4))); nil != err {
	//		return err.Relation(errors.ErrorServerInternalError("SHGWHE_LE.E_LE.U_TE.47"))
	//	}
	//
	//	return nil
	//}); nil != err {
	//	return response.ServerInternalError, nil, err
	//}

	if err := cls.example.UpdateExample(ctx, example, pointer.New(body.Field1), pointer.New(body.Field2), pointer.New(body.Field3), pointer.New(model.ExampleEnumFieldType(body.Field4))); nil != err {
		return response.ServerInternalError, nil, errors.NewI18n(i18n.CommonIdServerInternalError, err.Relation(errors.ErrorServerInternalError("SHGWHE_LE.E_LE.U_TE.563056")))
	}

	return response.Ok, nil, nil
}
