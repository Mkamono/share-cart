package shared

import "reflect"

func Cast[T any](v any) T {
	market := new(T)

	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)

	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		value := vv.Field(i).Interface()

		reflect.ValueOf(market).Elem().FieldByName(field.Name).Set(reflect.ValueOf(value))
	}

	return *market
}
