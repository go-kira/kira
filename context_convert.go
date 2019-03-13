package kira

import "strconv"

// ToString convert type f to string.
func (c *Context) ToString(f interface{}) string {
	switch f := f.(type) {
	case string:
		return f
	case bool:
		return strconv.FormatBool(f)
	case int:
		return strconv.Itoa(f)
	case int64:
		return strconv.FormatInt(f, 10)
	case float64:
		return strconv.FormatFloat(f, 'f', -1, 64)
	case nil:
	default:
		return ""
	}
	return ""
}

// ToInt convert type f to int.
func (c *Context) ToInt(f interface{}) (int, error) {
	switch f := f.(type) {
	case int:
		return f, nil
	case string:
		return strconv.Atoi(f)
	case bool:
		if f {
			return 1, nil
		}
		return 0, nil
	case int64:
		return int(f), nil
	case float64:
		return int(f), nil
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
	switch f := f.(type) {
	case bool:
		return f, nil
	case int:
		if f != 0 {
			return true, nil
		}
		return false, nil
	case int64:
		if f != 0 {
			return true, nil
		}
		return false, nil
	case float64:
		if f != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(f)
	}
	return false, nil
}

// ToFloat convert type f to float64.
func (c *Context) ToFloat(f interface{}) (float64, error) {
	switch f := f.(type) {
	case int:
		return float64(f), nil
	case int64:
		return float64(f), nil
	case string:
		return strconv.ParseFloat(f, 64)
	}
	return 0, nil
}
