package errors

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/i18n"
)

func NewI18n(id string, e Error, params ...string) I18nError {
	var p = make(map[string]string)

	for i := 0; i < len(params); i += 2 {
		if i+1 == len(params) {
			break
		}

		p[params[i]] = params[i+1]
	}

	return &i18e{id: id, e: e, params: p}
}

type I18nError interface {
	// I18n 处理I18N国际化消息
	I18n(ctx context.Context) string
}

type i18e struct {
	e      Error             // 基础错误信息
	id     string            // I18N国际化使用的消息ID
	params map[string]string // I18N国际化使用的命名参数表
}

func (cls *i18e) I18n(ctx context.Context) string {
	if nil == cls.params {
		cls.params = make(map[string]string)
	}

	cls.params["ErrorCodes"] = cls.e.Error()

	return i18n.New(cls.id, cls.params).Localize(ctx)
}
