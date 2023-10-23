package kitkit

import (
	"reflect"
)

func Allocate(val interface{}) error {
	rt := reflect.TypeOf(val)
	rv := reflect.ValueOf(val)

	if rt.Kind() != reflect.Ptr {
		return ErrorNotPointer
	}

	kind := rt.Elem().Kind()
	if kind == reflect.Func ||
		kind == reflect.Interface ||
		kind == reflect.UnsafePointer {
		return ErrorUnsupportedType
	}

	if rv.Elem().IsNil() {
		if rt.Elem().Kind() != reflect.Ptr {
			return ErrorNotPointer
		}
		rv.Elem().Set(reflect.New(rt.Elem().Elem()))
	}

	if kind == reflect.Ptr && rt.Elem().Elem().Kind() == reflect.Struct {
		for i := 0; i < rv.Elem().Elem().NumField(); i++ {
			field := rv.Elem().Elem().Field(i)
			kind := field.Type().Kind()
			if kind != reflect.Ptr {
				continue
			}

			_tmp := reflect.New(field.Type())
			err := Allocate(_tmp.Interface())
			if err != nil {
				return err
			}

			if field.CanSet() {
				field.Set(_tmp.Elem())
			}
		}
	}
	return nil
}
