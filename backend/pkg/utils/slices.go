package utils

func Contains[T comparable](element T, list []T) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}
	return false
}
