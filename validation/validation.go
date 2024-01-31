package validation

import (
	"golang.org/x/sync/errgroup"
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
