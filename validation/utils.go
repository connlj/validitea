package validation

import "reflect"

func unwrapReflectValue(v any) reflect.Value {
	switch v.(type) {
	case reflect.Value:
		return reflect.Indirect(v.(reflect.Value))
	default:
		return reflect.ValueOf(v)
	}
}
