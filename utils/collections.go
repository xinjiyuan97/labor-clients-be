package utils

func GetOrDefault[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func Intersect[T comparable](slice1 []T, slice2 []T) []T {
	result := make([]T, 0)
	for _, v := range slice1 {
		if Contains(slice2, v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T any, R any](slice []T, fn func(T) R) []R {
	result := make([]R, 0)
	for _, v := range slice {
		result = append(result, fn(v))
	}
	return result
}
