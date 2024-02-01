package validation

import (
	"golang.org/x/sync/errgroup"
	"reflect"
	"strings"
)

// Rule is the function type used by the validator to run against the given data
//
// Input:
//   - field: a string value representing the name of the value you are trying to validate ex. "username"
//   - value: the actual value this rule is going to be tested against
type Rule func(field string, value any) error

type pairing struct {
	data  any
	rules []Rule
}

type Validator struct {
	runners map[string]pairing
}

func New() Validator {
	validator := Validator{}
	validator.runners = make(map[string]pairing)

	return validator
}

func FromStructTags(structure any) Validator {
	validator := New()

	t := reflect.TypeOf(structure)
	v := reflect.ValueOf(structure)

	for i := 0; i < t.NumField() && t.Field(i).IsExported(); i++ {
		tagFull := t.Field(i).Tag.Get(TAG_VALIDATE)
		tags := strings.Split(tagFull, SEPARATOR)

		for _, tag := range tags {
			fieldName := t.Field(i).Name
			fieldValue := v.Field(i)

			switch {
			case strings.Contains(tag, SUBTAG_REQUIRED):
				//fmt.Printf("\tadding %v rule for value: %v\n", SUBTAG_REQUIRED, fieldName)
				validator.Add(fieldName, fieldValue, ValidateRequired)

			case strings.Contains(tag, SUBTAG_PRESENCE):
				//fmt.Printf("\tadding %v rule for value: %v\n", SUBTAG_PRESENCE, fieldName)
				validator.Add(fieldName, fieldValue, ValidatePresence)

			case strings.Contains(tag, SUBTAG_MIN):
				//fmt.Printf("\tadding %v rule for value: %v\n", SUBTAG_MIN, fieldName)
				validator.Add(fieldName, fieldValue, ValidateMinLength(2))

			case strings.Contains(tag, SUBTAG_MAX):
				//fmt.Printf("\tadding %v rule for value: %v\n", SUBTAG_MAX, fieldName)
				validator.Add(fieldName, fieldValue, ValidateMaxLength(32))

			case strings.Contains(tag, SUBTAG_EMAIL):
				//fmt.Printf("\tadding %v rule for value: %v\n", SUBTAG_EMAIL, fieldName)
			default:
				//fmt.Printf("unknown rule request for value: %v\n", fieldName)
			}
		}
	}

	return validator
}

func (v *Validator) Add(field string, data any, rules ...Rule) {
	val, ok := v.runners[field]

	if ok {
		val.rules = append(val.rules, rules...)
		v.runners[field] = val
	}

	if !ok {
		v.runners[field] = pairing{data, rules}
	}
}

func (v *Validator) Validate() error {
	eg := errgroup.Group{}

	for field, runner := range v.runners {
		func(field string, runner pairing) {
			eg.Go(func() error {
				for _, rule := range runner.rules {
					if err := rule(field, runner.data); err != nil {
						return err
					}
				}
				return nil
			})
		}(field, runner)
	}

	return eg.Wait()
}
