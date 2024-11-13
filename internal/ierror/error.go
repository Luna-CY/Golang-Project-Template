package ierror

import (
	"errors"
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/internal/i18n"
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
)

var ErrorRecordNotFound = I18n("common-record-not-found")
var ErrorServerInternalError = I18n("common-server-internal-error")

func New(message string, params ...any) *Error {
	return &Error{message: message, values: params}
}

func I18n(id string, params ...string) *Error {
	var p = make(map[string]string)

	for i := 0; i < len(params); i += 2 {
		if i+1 == len(params) {
			break
		}

		p[params[i]] = params[i+1]
	}

	return &Error{id: id, params: p}
}

type Error struct {
	id      string
	message string
	values  []any
	params  map[string]string
}

func (e *Error) Error() string {
	return fmt.Sprintf(e.message, e.values...)
}

func (e *Error) Is(target error) bool {
	var ie *Error

	if errors.As(target, &ie) {
		if nil != ie && ie.id == e.id && ie.id != "" {
			return true
		}
	}

	return e.Error() == target.Error()
}

func (e *Error) I18n(ctx icontext.Context) string {
	return i18n.New(e.id, e.params).Localize(ctx)
}
