package helpers

// Difference -  returns the values in slice1 that are not present in any of the other slices.
func Difference(slice []string, slice2 []string) []string {
	var result []string

	// range over slice
	for _, v := range slice {
		if !Contains(slice2, v) {
			result = append(result, v)
		}
	}

	return result
}
