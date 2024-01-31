package validation

import (
	"fmt"
	"reflect"
)

func ValidateRequired(field string, value any) error {
	if value == nil {
		return fmt.Errorf("%s can't be nil", field)
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String:
		if val.Len() == 0 || value.(string) == "" {
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

	fmt.Printf("validated required of %v\n", field)
	return nil
}

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

	fmt.Printf("validated presence of %v\n", field)
	return nil
}

func ValidateMinLength(min int) Rule {
	return func(field string, value any) error {
		fmt.Println(field, value, reflect.ValueOf(value), reflect.ValueOf(value).Kind())

		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
			fmt.Println("in case")
			if val.Len() < min {
				return fmt.Errorf("value { %+v } is shorter than the minimum of %+v", val, min)
			}
		default:
			fmt.Println("default case")
		}

		fmt.Printf("validated min length of %v\n", field)
		return nil
	}
}

func ValidateMaxLength(max int) Rule {
	return func(field string, value any) error {

		fmt.Printf("validated max length of %v\n", field)
		return nil
	}
}
