package array

func Diff[T comparable](a, b []T) []T {
	var c []T
	for _, t := range a {
		if Search(t, b) < 0 {
			c = append(c, t)
		}
	}
	return c
}
