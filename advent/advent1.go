package advent

import (
	"github.com/nveeser/go-experiment/xiter"
	"iter"
	"sort"
)

func Part1(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	var got int
	for l, r := range xiter.ZipV(asIter(left), asIter(right)) {
		if l.V > r.V {
			got += l.V - r.V
		} else {
			got += r.V - l.V
		}
		//fmt.Printf("%d %d => %d\n", l.V, r.V, got)
	}
	return got
}

func Part2(left, right []int) int {
	counts := make(map[int]int)
	for _, r := range right {
		x := counts[r]
		counts[r] = x + 1
		//fmt.Printf("%d => %d\n", r, counts[r])
	}
	var val int
	for _, l := range left {
		n := counts[l]
		val += l * n
		//fmt.Printf("%d (%d) => %d\n", l, n, val)
	}
	return val
}

func asIter[T any](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}
