package _map

func Values[T comparable, V any](m map[T]V) []V {
	var t = []V{}
	for _, v := range m {
		t = append(t, v)
	}
	return t
}
