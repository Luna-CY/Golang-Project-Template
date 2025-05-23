package i18n

import (
	"strings"
	"text/template"

	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/language"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
)

func New(id string, params map[string]string) Localize {
	return Localize{id: id, params: params}
}

type Localize struct {
	id     string
	params map[string]string
}

// Localize 本地化
func (cls Localize) Localize(ctx context.Context) string {
	var messages = pointer.Or(languages[GetAcceptLanguage(ctx)], languages[language.SimpleChinese])

	message, ok := messages[cls.id]
	if !ok {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I18N: not found ID: %s", cls.id)

		return ""
	}

	tp, err := template.New(cls.id).Parse(message)
	if nil != err {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I18N: parse I18N template failed: %s", err)

		return ""
	}

	var buffer = new(strings.Builder)
	if err := tp.Execute(buffer, cls.params); nil != err {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I18N: template processing failed: %s", err)

		return ""
	}

	return buffer.String()
}

func GetAcceptLanguage(ctx context.Context) language.Language {
	var value = ctx.Value("accept-language")
	if nil != value {
		return value.(language.Language)
	}

	return language.SimpleChinese
}
