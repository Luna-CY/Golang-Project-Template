package itype

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
)

type MultiValues[T any] []T

func (cls *MultiValues[T]) Scan(value any) error {
	if s, ok := value.(string); ok {
		if "" == s {
			*cls = make(MultiValues[T], 0)

			return nil
		}

		if err := json.Unmarshal([]byte(s), cls); nil != err {
			return err
		}

		return nil
	}

	if bs, ok := value.([]byte); ok {
		if 0 == len(bs) {
			*cls = make(MultiValues[T], 0)

			return nil
		}

		if err := json.Unmarshal(bs, cls); nil != err {
			return err
		}

		return nil
	}

	return errors.New(errors.ErrorTypeServerInternalError, "invalid values")
}

func (cls *MultiValues[T]) Value() (driver.Value, error) {
	if nil == cls || 0 == len(*cls) {
		return "[]", nil
	}

	var value, err = json.Marshal(cls)
	if nil != err {
		return nil, err
	}

	return string(value), nil
}
