package _map

func Merge[k comparable, v any](maps ...map[k]v) map[k]v {
	c := make(map[k]v)
	for _, mapItem := range maps {
		for ki, vi := range mapItem {
			c[ki] = vi
		}
	}
	return c
}
