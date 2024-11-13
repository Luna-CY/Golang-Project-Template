package pointer

import "reflect"

func New[T any](v T) *T {
	return &v
}

func Or[T any](v T, d T) T {
	if reflect.ValueOf(v).IsZero() {
		return d
	}

	return v
}

func NewOrNil[T any](v T) *T {
	if reflect.ValueOf(v).IsZero() {
		return nil
	}

	return &v
}
