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
func (c *Context) ToInt(f interface{}) int {
	switch f.(type) {
	case int:
		return f.(int)
	case string:
		in, err := strconv.Atoi(f.(string))
		if err != nil {
			c.Error(err)
		}
		return in
	case bool:
		if f.(bool) {
			return 1
		}
		return 0
	case int64:
		return int(f.(int64))
	case float64:
		return int(f.(float64))
	case nil:
	default:
		return 0
	}
	return 0
}

// ToInt64 convert type f to int64.
func (c *Context) ToInt64(f interface{}) int64 {
	return int64(c.ToInt(f))
}

// ToBool convert type f to bool.
func (c *Context) ToBool(f interface{}) bool {
	switch f.(type) {
	case bool:
		return f.(bool)
	case int:
		if f.(int) != 0 {
			return true
		}
		return false
	case int64:
		if f.(int64) != 0 {
			return true
		}
		return false
	case float64:
		if f.(float64) != 0 {
			return true
		}
		return false
	case string:
		bol, err := strconv.ParseBool(f.(string))
		if err != nil {
			c.Error(err)
		}
		return bol
	}
	return false
}

// ToFloat convert type f to float64.
func (c *Context) ToFloat(f interface{}) float64 {
	switch f.(type) {
	case int:
		return float64(f.(int))
	case int64:
		return float64(f.(int64))
	case string:
		bol, err := strconv.ParseFloat(f.(string), 64)
		if err != nil {
			c.Error(err)
		}
		return bol
	}
	return 0
}
