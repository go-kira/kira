package kira

import "strconv"

// ToString convert type f to string.
func (c *Context) ToString(f interface{}) string {
	switch f.(type) {
	case string:
		return f.(string)
	case bool:
		return strconv.FormatBool(f.(bool))
	case int:
		return strconv.Itoa(f.(int))
	case int64:
		return strconv.FormatInt(f.(int64), 10)
	case float64:
		return strconv.FormatFloat(f.(float64), 'f', -1, 64)
	case nil:
	default:
		return ""
	}
	return ""
}

// ToInt convert type f to int.
func (c *Context) ToInt(f interface{}) (int, error) {
	switch f.(type) {
	case int:
		return f.(int), nil
	case string:
		return strconv.Atoi(f.(string))
	case bool:
		if f.(bool) {
			return 1, nil
		}
		return 0, nil
	case int64:
		return int(f.(int64)), nil
	case float64:
		return int(f.(float64)), nil
	case nil:
	default:
		return 0, nil
	}
	return 0, nil
}

// ToInt64 convert type f to int64.
func (c *Context) ToInt64(f interface{}) (int64, error) {
	in, err := c.ToInt(f)
	if err != nil {
		return 0, err
	}

	return int64(in), nil
}

// ToBool convert type f to bool.
func (c *Context) ToBool(f interface{}) (bool, error) {
	switch f.(type) {
	case bool:
		return f.(bool), nil
	case int:
		if f.(int) != 0 {
			return true, nil
		}
		return false, nil
	case int64:
		if f.(int64) != 0 {
			return true, nil
		}
		return false, nil
	case float64:
		if f.(float64) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(f.(string))
	}
	return false, nil
}

// ToFloat convert type f to float64.
func (c *Context) ToFloat(f interface{}) (float64, error) {
	switch f.(type) {
	case int:
		return float64(f.(int)), nil
	case int64:
		return float64(f.(int64)), nil
	case string:
		return strconv.ParseFloat(f.(string), 64)
	}
	return 0, nil
}
