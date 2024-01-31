package rules

import (
	"fmt"
	"reflect"
)

func ValidatePresence(field string, value any) error {
	if value == nil {
		return fmt.Errorf("%s can't be nil", field)
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String:
		if val.Len() == 0 {
			return fmt.Errorf("%s can't be blank", field)
		}
	case reflect.Array, reflect.Slice, reflect.Map:
		if val.Len() == 0 {
			return fmt.Errorf("%s can't be blank", field)
		}
	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return fmt.Errorf("%s can't be nil", field)
		}
	}

	fmt.Printf("Presence of %s validated\n", field)
	return nil
}
