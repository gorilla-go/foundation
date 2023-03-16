package array

func Unique[T comparable](arr []T) []T {
	var m = make(map[T]int8)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = 0
	}
	var f []T
	for k, _ := range m {
		f = append(f, k)
	}
	return f
}
