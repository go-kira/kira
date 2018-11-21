package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-kira/kira/helpers"
)

// Errors
var (
	errorNotString = errors.New("")
)

// Validation struct
type Validation struct {
	Separator       string
	ParamsPrefix    string
	ParamsSeparator string
}

// New return Validation instance
func New() *Validation {
	return &Validation{
		Separator:       "|",
		ParamsPrefix:    ":",
		ParamsSeparator: ",",
	}
}

// Validate for check the validation
func (v *Validation) Validate(field string, value interface{}, rules string) error {
	split := strings.Split(rules, v.Separator)

	for _, rule := range split {
		// get before : in the rule
		prefix := helpers.Before(rule, ":")
		if prefix == "" {
			prefix = rule
		}
		// rule params if exists
		suffix := helpers.After(rule, ":")
		params := strings.Split(suffix, v.ParamsSeparator)
		// nil the params if there no suffix
		if suffix == "" {
			params = nil
		}

		if ruleValue, ruleInMap := RulesNames[prefix]; ruleInMap {
			// do the validation
			if checkErr := ruleValue(field, value, params...); checkErr != nil {
				return checkErr
			}
		} else {
			return fmt.Errorf("the `%s` rule not supported", rule)
		}

	}

	// validation pass
	return nil
}
