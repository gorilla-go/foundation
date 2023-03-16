package array

func Walk[T any](a []T, f func(*T)) {
	for i := 0; i < len(a); i++ {
		f(&a[i])
	}
}
