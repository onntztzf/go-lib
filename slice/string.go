package slice

// ContainsString checks if a string exists in a list of strings.
func ContainsString(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}
	return false
}

// RemoveDuplicateStrings removes duplicate strings from a string slice.
func RemoveDuplicateStrings(arr []string) []string {
	occurred := map[string]bool{}
	result := make([]string, 0, len(arr))
	for _, value := range arr {
		if !occurred[value] {
			occurred[value] = true
			result = append(result, value)
		}
	}
	return result
}

// RemoveStringElements removes specific string elements from a string slice.
func RemoveStringElements(slice []string, targets ...string) []string {
	for _, r := range targets {
		for i := 0; i < len(slice); i++ {
			if slice[i] == r {
				slice = append(slice[:i], slice[i+1:]...)
				i--
			}
		}
	}
	return slice
}
