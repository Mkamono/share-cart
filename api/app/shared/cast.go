package shared

import (
	"reflect"
)

// Projection 引数の構造体を型引数に指定した構造体に変換する
// 同じ名前のフィールドがある場合、型が一致している場合のみ値をコピーする
func Projection[T any](v any) T {
	var market T
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	mt := reflect.TypeOf(market)
	mv := reflect.ValueOf(&market).Elem()

	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		value := vv.Field(i)

		if mf, ok := mt.FieldByName(field.Name); ok {
			if value.Type().AssignableTo(mf.Type) {
				mv.FieldByName(field.Name).Set(value)
			}
		}
	}

	return market
}
