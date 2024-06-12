package utils

// Adds an element to the beginning of a slice
func Prepend[T any](slice []T, value T) []T {
	return append([]T{value}, slice...)
}
