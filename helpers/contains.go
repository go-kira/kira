package helpers

// Contains - check if the slice of strings containt the given string
func Contains(vals []string, s string) bool {
	for _, v := range vals {
		if v == s {
			return true
		}
	}

	return false
}
