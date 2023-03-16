package _map

func Keys[T comparable, S any](m map[T]S) []T {
	var arr []T
	for k, _ := range m {
		arr = append(arr, k)
	}
	return arr
}
