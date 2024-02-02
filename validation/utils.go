package validation

import (
	"reflect"
)

func unwrapReflectValue(v any) reflect.Value {
	switch v.(type) {
	case reflect.Value:
		return reflect.Indirect(v.(reflect.Value))
	default:
		reflectValue := reflect.Indirect(reflect.ValueOf(v))
		for reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
			reflectValue = reflect.Indirect(reflectValue)
		}
		return reflectValue

	}
}
