package errors

import (
	"errors"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/i18n"
	"strings"
)

func ErrorRecordNotFound(code string) Error {
	return I18n(ErrorTypeRecordNotFound, code, "common-record-not-found")
}

func ErrorServerInvalidRequest(code string) Error {
	return I18n(ErrorTypeInvalidRequest, code, "common-invalid-request")
}

func ErrorServerInternalError(code string) Error {
	return I18n(ErrorTypeServerInternalError, code, "common-server-internal-error")
}

type ErrorType int

const (
	ErrorTypeRecordNotFound = ErrorType(iota)
	ErrorTypeInvalidRequest
	ErrorTypeServerInternalError
)

type Error interface {
	error

	// IsType 检查给定的错误是否是某个类型的错误
	IsType(ErrorType) bool

	// Relation 将新的错误关联到当前错误
	// 该方法返回当前错误对象
	Relation(...Error) Error

	// I18n 处理I18N国际化消息
	I18n(ctx context.Context) string
}

func New(t ErrorType, code string, message any, params ...any) Error {
	return &IError{t: t, code: code, message: message, values: params}
}

func I18n(t ErrorType, id string, code string, params ...string) *IError {
	var p = make(map[string]string)

	for i := 0; i < len(params); i += 2 {
		if i+1 == len(params) {
			break
		}

		p[params[i]] = params[i+1]
	}

	return &IError{id: id, t: t, code: code, params: p}
}

type IError struct {
	id        string            // I18N国际化使用的消息ID
	t         ErrorType         // 错误类型
	code      string            // 具体的错误编码，硬编码在发生错误的地方
	message   any               // 消息内容，可以为任何支持%v显示的内容
	values    []any             // 变量列表，如果此列表不为空，message必须是一个可以提供给 fmt.Sprintf 方法的模板字符串
	params    map[string]string // I18N国际化使用的命名参数表
	relations []Error           // 关联的错误列表
}

func (cls *IError) Error() string {
	var sb = new(strings.Builder)
	if 0 != len(cls.values) {
		sb.WriteString(fmt.Sprintf("%s: %s", cls.code, fmt.Sprintf(cls.message.(string), cls.values...)))
	} else {
		sb.WriteString(fmt.Sprintf("%v", cls.message))
	}

	if 0 != len(cls.relations) {
		sb.WriteString(" (")

		for _, r := range cls.relations {
			sb.WriteString(r.Error())
		}

		sb.WriteString(")")
	}

	return sb.String()
}

func (cls *IError) IsType(errorType ErrorType) bool {
	return cls.t == errorType
}

func (cls *IError) Relation(errs ...Error) Error {
	cls.relations = append(cls.relations, errs...)

	return cls
}

func (cls *IError) I18n(ctx context.Context) string {
	if "" == cls.id {
		return cls.Error()
	}

	return i18n.New(cls.id, cls.params).Localize(ctx)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
