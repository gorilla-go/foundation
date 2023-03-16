package array

func Reduce[T any, M any](i []T, f func(M, T) M, n M) M {
	acc := n
	for _, item := range i {
		acc = f(acc, item)
	}
	return acc
}
