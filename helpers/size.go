package helpers

import (
	"unicode/utf8"
)

// Size return size of string or float64 or float32 or slice
func Size(value interface{}) float64 {
	switch value.(type) {
	// convert string value to integer
	case string:
		return float64(utf8.RuneCountInString(value.(string)))
	case float32:
		return float64(value.(float32))
	case int:
		return float64(value.(int))
	case []string:
		return float64(len(value.([]string)))
	case []int:
		return float64(len(value.([]int)))
	case []float32:
		return float64(len(value.([]float32)))
	case []float64:
		return float64(len(value.([]float64)))
	default:
		return value.(float64)
	}
}
