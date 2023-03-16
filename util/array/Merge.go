package array

func Merge[T any](a, b []T) []T {
	a = append(a, b...)
	return a
}
