package validation

import (
	"fmt"
	neturl "net/url"
	"strconv"
	"strings"

	"github.com/go-kira/kira/helpers"
)

// FuncString this for all rules that accept string as value
type FuncString func(field string, value interface{}, parameters ...string) error

// RulesNames ...
var RulesNames = map[string]FuncString{
	"required": required,
	"integer":  integer,
	"numeric":  numeric,
	"max":      max,
	"min":      min,
	"between":  between,
	"email":    email,
	"url":      url,
}

// Rules ...
func required(field string, value interface{}, parameters ...string) error {
	if len(strings.TrimSpace(helpers.ConvertToString(value))) == 0 {
		return fmt.Errorf(errRequired, field)
	}
	// not empty
	return nil
}
func integer(field string, value interface{}, parameters ...string) error {
	_, err := strconv.Atoi(helpers.ConvertToString(value))
	if err != nil {
		return fmt.Errorf(errInteger, field)
	}

	// not empty
	return nil
}
func numeric(field string, value interface{}, parameters ...string) error {
	_, err := strconv.ParseFloat(helpers.ConvertToString(value), 64)
	if err != nil {
		return fmt.Errorf(errNumeric, field)
	}

	// not empty
	return nil
}

func max(field string, value interface{}, parameters ...string) error {
	param := helpers.ConvertToFloat64(parameters[0]) // converted to float64

	if size := helpers.Size(value); size > param {
		return fmt.Errorf(errMax, field, param)
	}

	// not empty
	return nil
}

func min(field string, value interface{}, parameters ...string) error {
	param := helpers.ConvertToFloat64(parameters[0]) // converted to float64

	if size := helpers.Size(value); size < param {
		return fmt.Errorf(errMin, field, param)
	}

	// not empty
	return nil
}

func between(field string, value interface{}, parameters ...string) error {
	param1 := helpers.ConvertToFloat64(parameters[0]) // converted to float64
	param2 := helpers.ConvertToFloat64(parameters[1]) // converted to float64

	size := helpers.Size(value)
	if size < param1 || size > param2 {
		return fmt.Errorf(errBetween, size, param1, param2)
	}

	// not empty
	return nil
}
func email(field string, value interface{}, parameters ...string) error {

	if isEmail := rxEmail.MatchString(helpers.ConvertToString(value)); isEmail == false {
		return fmt.Errorf(errEmail, field)
	}

	// not empty
	return nil
}

func url(field string, value interface{}, parameters ...string) error {
	if helpers.ConvertToString(value) == "" {
		return nil
	}

	if _, err := neturl.ParseRequestURI(helpers.ConvertToString(value)); err != nil {
		return fmt.Errorf(errURL, field)
	}

	// not empty
	return nil
}
