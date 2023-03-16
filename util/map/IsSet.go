package _map

func IsSet[k comparable, v any](key k, m map[k]v) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}
