package array

func Search[T comparable](s T, a []T) int {
	for i := 0; i < len(a); i++ {
		if s == a[i] {
			return i
		}
	}
	return -1
}
