package errors

import (
	"fmt"
	"strings"
)

func ErrorRecordNotFound(code string) Error {
	return New(ErrorTypeRecordNotFound, code, "数据不存在")
}

func ErrorInvalidRequest(code string) Error {
	return New(ErrorTypeInvalidRequest, code, "参数错误")
}

func ErrorServerInternalError(code string) Error {
	return New(ErrorTypeServerInternalError, code, "服务器内部错误")
}

type ErrorType int

const (
	ErrorTypeRecordNotFound = ErrorType(iota)
	ErrorTypeInvalidRequest
	ErrorTypeServerInternalError
)

type Error interface {
	error

	// GetCode 获取错误代码
	GetCode() string

	// IsType 检查给定的错误是否是某个类型的错误
	IsType(ErrorType) bool

	// Relation 将新的错误关联到当前错误
	// 该方法返回当前错误对象
	Relation(...Error) Error

	// Relations 获取所有关联的错误列表
	Relations() []Error
}

func New(t ErrorType, code string, message any, params ...any) Error {
	return &IError{t: t, code: code, message: message, values: params}
}

type IError struct {
	t         ErrorType // 错误类型
	code      string    // 具体的错误编码，硬编码在发生错误的地方
	message   any       // 消息内容，可以为任何支持%v显示的内容
	values    []any     // 变量列表，如果此列表不为空，message必须是一个可以提供给 fmt.Sprintf 方法的模板字符串
	relations []Error   // 关联的错误列表
}

func (cls *IError) Error() string {
	var sb = new(strings.Builder)

	if 0 != len(cls.values) {
		sb.WriteString(fmt.Sprintf("%s: %s", cls.code, fmt.Sprintf(cls.message.(string), cls.values...)))
	} else {
		sb.WriteString(fmt.Sprintf("%s: %v", cls.code, cls.message))
	}

	var es []string
	for _, relation := range cls.relations {
		es = append(es, relation.Error())
	}

	if 0 != len(es) {
		sb.WriteString("(" + strings.Join(es, ",") + ")")
	}

	return sb.String()
}

func (cls *IError) GetCode() string {
	return cls.code
}

func (cls *IError) IsType(errorType ErrorType) bool {
	return cls.t == errorType
}

func (cls *IError) Relation(errs ...Error) Error {
	cls.relations = append(cls.relations, errs...)

	return cls
}

func (cls *IError) Relations() []Error {
	return cls.relations
}
