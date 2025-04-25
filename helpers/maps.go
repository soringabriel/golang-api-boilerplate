package helpers

import (
	"fmt"
	"reflect"
)

func StructToMapString(s any) map[string]string {
	out := make(map[string]string)
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		key := fieldType.Name

		if !field.IsValid() || (field.Kind() == reflect.Ptr && field.IsNil()) {
			continue // skip nil pointers
		}

		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		out[key] = fmt.Sprintf("%v", field.Interface())
	}
	return out
}
