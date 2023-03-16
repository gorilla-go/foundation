package array

import (
	"sort"
)

type SortMode int

const DESC SortMode = 1
const ASC SortMode = 0

func Sort[T any, Y int](a []T, b []int, mode SortMode) []T {
	if len(a) != len(b) {
		panic("invalid array lens.")
	}

	var bc = make(map[int]int)
	for i := 0; i < len(b); i++ {
		bc[i] = b[i]
	}

	if mode == ASC {
		sort.Ints(b)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(b)))
	}

	var bcArr []int
	for i := 0; i < len(b); i++ {
		for seq, k := range bc {
			if k == b[i] {
				bcArr = append(bcArr, seq)
				delete(bc, seq)
			}
		}
	}
	var r []T
	for _, v := range bcArr {
		r = append(r, a[v])
	}

	return r
}
