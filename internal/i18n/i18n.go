package i18n

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/language"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"strings"
	"text/template"
)

func New(id string, params map[string]string) Localize {
	return Localize{id: id, params: params}
}

type Localize struct {
	id     string
	params map[string]string
}

// Localize 本地化
func (cls Localize) Localize(ctx icontext.Context) string {
	var lang = string(GetAcceptLanguage(ctx))

	messages, ok := languages[lang]
	if !ok {
		logger.SugarLogger(ctx).Errorf("I18N: not found language configuration: %s", lang)

		return ""
	}

	message, ok := messages[cls.id]
	if !ok {
		logger.SugarLogger(ctx).Errorf("I18N: not found ID: %s", cls.id)

		return ""
	}

	tp, err := template.New(cls.id).Parse(message)
	if nil != err {
		logger.SugarLogger(ctx).Errorf("I18N: parse I18N template failed: %s", err)

		return ""
	}

	var buffer = new(strings.Builder)
	if err := tp.Execute(buffer, cls.params); nil != err {
		logger.SugarLogger(ctx).Errorf("I18N: template processing failed: %s", err)

		return ""
	}

	return buffer.String()
}

func GetAcceptLanguage(ctx icontext.Context) language.Language {
	var value = ctx.Value("accept-language")
	if nil != value {
		return value.(language.Language)
	}

	return language.SimpleChinese
}
