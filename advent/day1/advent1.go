package day1

import (
	"github.com/nveeser/go-experiment/xiter"
	"iter"
	"log"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

func Read(data string) (l, r []int) {
	for _, row := range strings.Split(data, "\n") {
		if strings.TrimFunc(row, unicode.IsSpace) == "" {
			continue
		}
		parts := strings.Fields(row)
		if len(parts) != 2 {
			log.Fatalf("Bad Parts: %s", parts)
		}
		x, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatalf("Error parsing %s: %v", parts[0], err)
		}
		y, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatalf("Error parsing %s: %v", parts[1], err)
		}
		l = append(l, int(x))
		r = append(r, int(y))
	}
	return
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
