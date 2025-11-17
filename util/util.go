package util

func FirstOrDefault[T any](slice []T, defaultValue T) T {
	if len(slice) > 0 {
		return slice[0]
	}
	return defaultValue
}
