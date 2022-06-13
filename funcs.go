package utils

// If generate quick if
func If[T any](cond bool, yes T, no T) T {
	if cond {
		return yes
	}
	return no
}
