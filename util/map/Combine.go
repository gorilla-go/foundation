package _map

func Combine[K comparable, V any](a []K, b []V) map[K]V {
	if len(a) != len(b) {
		panic("invalid array lens.")
	}
	var c = make(map[K]V)
	for i := 0; i < len(a); i++ {
		c[a[i]] = b[i]
	}
	return c
}
