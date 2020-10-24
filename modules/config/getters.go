package config

// GetDefault return value by key, if the value empty set a default value.
func (k *Config) GetDefault(key string, def interface{}) interface{} {
	k.lock.RLock()
	defer k.lock.RUnlock()

	val := k.Get(key)
	if val == nil {
		return def
	}

	return val
}

// GetString - return config as string type.
func (k *Config) GetString(key string, def ...string) string {
	if k.Get(key) == nil {
		if len(def) > 0 {
			return def[0]
		}

		return ""
	}

	return k.Get(key).(string)
}

// GetInt - return config as int type.
func (k *Config) GetInt(key string, def ...int) int {
	if k.Get(key) == nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}

	return int(k.Get(key).(int64))
}

// GetInt64 - return config as int type.
func (k *Config) GetInt64(key string, def ...int) int64 {
	return int64(k.GetInt(key, def...))
}

// GetFloat64 - return config as float64 type.
func (k *Config) GetFloat64(key string, def ...float64) float64 {
	if k.Get(key) == nil {
		if len(def) > 0 {
			return def[0]
		}
		return 0
	}

	return k.Get(key).(float64)
}

// GetBool - return config as int type.
func (k *Config) GetBool(key string, def ...bool) bool {
	if k.Get(key) == nil {
		if len(def) > 0 {
			return def[0]
		}
		return false
	}

	return k.Get(key).(bool)
}

// // GetArray return a slice of empty interfaces.
// func (k *Config) GetArray(key string) []interface{} {
// 	return k.Get(key).([]interface{})
// }
//
// // GetArrayString return a slice of string.
// func (k *Config) GetArrayString(key string) (slice []string) {
// 	if k.Get(key) != nil {
// 		for _, v := range k.Get(key).([]interface{}) {
// 			slice = append(slice, v.(string))
// 		}
// 	}
//
// 	return slice
// }
//
// // GetArrayInt return a slice of int.
// func (k *Config) GetArrayInt(key string) (slice []int) {
// 	if k.Get(key) != nil {
// 		for _, v := range k.Get(key).([]interface{}) {
// 			slice = append(slice, int(v.(int64)))
// 		}
// 	}
// 	return slice
// }
//
// // GetArrayInt64 return a slice of int64.
// func (k *Config) GetArrayInt64(key string) (slice []int64) {
// 	if k.Get(key) != nil {
// 		for _, v := range k.Get(key).([]interface{}) {
// 			slice = append(slice, v.(int64))
// 		}
// 	}
// 	return slice
// }
