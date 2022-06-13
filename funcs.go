package utils

// If generate quick if
func If[T any](cond bool, yes T, no T) T {
	if cond {
		return yes
	}
	return no
}

// Contains check if slice contains  item
func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
