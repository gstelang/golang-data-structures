package common

func Includes[T comparable](arr []T, e T) bool {
	for _, elem := range arr {
		if elem == e {
			return true
		}
	}
	return false
}
