package pointer

import (
	"reflect"
)

func Default[T any](v T) T {
	var vv = reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	var vt = reflect.TypeOf(v)
	if reflect.Ptr == vt.Kind() {
		vt = vt.Elem()
	}

	switch vt.Kind() {
	case reflect.Struct:
		var cv = reflect.New(vt).Elem()

		for i := 0; i < cv.NumField(); i++ {
			var cvf = cv.Field(i)

			if reflect.Ptr == cvf.Kind() && cvf.IsZero() {
				continue
			}

			cvf.Set(reflect.ValueOf(Default(vv.Field(i).Interface())))
		}

		return cv.Interface().(T)
	case reflect.Slice:
		if !vv.IsZero() && 0 != vv.Len() {
			return v
		}

		return reflect.MakeSlice(vt, 0, 0).Interface().(T)
	case reflect.Map:
		if !vv.IsZero() && 0 != vv.Len() {
			return v
		}

		return reflect.MakeMap(vt).Interface().(T)
	default:
		if !vv.IsZero() {
			return v
		}

		return reflect.New(vt).Elem().Interface().(T)
	}
}
