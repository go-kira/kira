package helpers

import (
	"log"
	"strconv"
)

// ConvertToString convert any data type to specific type
func ConvertToString(value interface{}) string {
	switch value.(type) {
	case bool:
		return strconv.FormatBool(value.(bool))
	case int:
		return strconv.Itoa(value.(int))
	case float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case nil:
		return ""
	default:
		return value.(string)
	}
}

// ConvertToInteger convert any data type to specific type
func ConvertToInteger(value interface{}) int {
	switch value.(type) {
	case string:
		if value.(string) == "" {
			return 0
		}

		integer, err := strconv.Atoi(value.(string))
		if err != nil {
			log.Fatal(err)
		}
		return integer
	case nil:
		return 0
	default:
		return value.(int)
	}
}

// ConvertToFloat64 convert any data type to float64
func ConvertToFloat64(value interface{}) float64 {
	switch value.(type) {
	case string:
		float, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			log.Fatal(err)
		}
		return float
	case int:
		return float64(value.(float64))
	default:
		return value.(float64)
	}
}

// ConvertToBool convert any data type to bool
func ConvertToBool(value interface{}) bool {
	switch value.(type) {
	case string:
		bool, err := strconv.ParseBool(value.(string))
		if err != nil {
			log.Fatal(err)
		}
		return bool
	default:
		return value.(bool)
	}
}
