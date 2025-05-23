package request

import (
	"reflect"
	"strings"

	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/gin-gonic/gin"
)

type Handler func(dst any) errors.Error

func BindHandlerTrimSliceEmptyValue(dst any) errors.Error {
	var rv = reflect.ValueOf(dst)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			if err := BindHandlerTrimSliceEmptyValue(rv.Field(i).Addr().Interface()); nil != err {
				return err.Relation(errors.ErrorServerInternalError("SHR_ST.BHTSEV_UE.233330"))
			}
		}
	case reflect.Slice:
		if 0 == rv.Len() || rv.Index(0).Kind() != reflect.String {
			return nil
		}

		var values []string

		for i := 0; i < rv.Len(); i++ {
			if "" == rv.Index(i).String() {
				continue
			}

			values = append(values, rv.Index(i).String())
		}

		rv.Set(reflect.ValueOf(values))
	default:
	}

	return nil
}

func ShouldBindJSON(c *gin.Context, dst any, handlers ...Handler) errors.Error {
	if err := c.ShouldBindJSON(dst); nil != err {
		return errors.New(errors.ErrorTypeInvalidRequest, "SHR_ST.SBJ_ON.503337", err)
	}

	if err := trimStringSpace(dst); nil != err {
		return err.Relation(errors.ErrorServerInternalError("SHR_ST.SBJ_ON.543340"))
	}

	for _, handler := range handlers {
		if err := handler(dst); nil != err {
			return err.Relation(errors.ErrorServerInternalError("SHR_ST.SBJ_ON.593345"))
		}
	}

	return nil
}

func ShouldBindForm(c *gin.Context, dst any, handlers ...Handler) errors.Error {
	if err := c.ShouldBind(dst); nil != err {
		return errors.New(errors.ErrorTypeInvalidRequest, "SHR_ST.SBF_RM.683351", err)
	}

	if err := trimStringSpace(dst); nil != err {
		return err.Relation(errors.ErrorServerInternalError("SHR_ST.SBF_RM.723354"))
	}

	for _, handler := range handlers {
		if err := handler(dst); nil != err {
			return err.Relation(errors.ErrorServerInternalError("SHR_ST.SBF_RM.773401"))
		}
	}

	return nil
}

func trimStringSpace(v any) errors.Error {
	var rv = reflect.ValueOf(v)

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			var field = rv.Field(i)
			if reflect.Ptr == field.Kind() {
				field = field.Elem()
			}

			if reflect.Interface == field.Kind() {
				field = field.Elem()
			}

			if err := trimStringSpace(field.Addr().Interface()); nil != err {
				return err.Relation(errors.ErrorServerInternalError("SHR_ST.TSS_CE.1043407"))
			}
		}
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			var item = rv.Index(i)
			if reflect.Ptr == item.Kind() {
				item = item.Elem()
			}

			if reflect.Interface == item.Kind() {
				item = item.Elem()
			}

			if reflect.String == item.Kind() {
				item.SetString(strings.TrimSpace(item.String()))
			}
		}
	case reflect.Map:
		for _, key := range rv.MapKeys() {
			var value = rv.MapIndex(key)
			if reflect.Ptr == value.Kind() {
				value = value.Elem()
			}

			if reflect.Interface == value.Kind() {
				value = value.Elem()
			}

			switch value.Kind() {
			case reflect.String:
				rv.SetMapIndex(key, reflect.ValueOf(strings.TrimSpace(value.String())))
			case reflect.Struct, reflect.Slice, reflect.Map:
				if err := trimStringSpace(value.Interface()); nil != err {
					return err.Relation(errors.ErrorServerInternalError("SHR_ST.TSS_CE.1383413"))
				}
			default:
				// 其他类型不处理
			}
		}
	case reflect.String:
		rv.SetString(strings.TrimSpace(rv.String()))
	default:
		// 其他类型不处理
	}

	return nil
}
