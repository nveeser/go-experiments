package bitbit

import (
	"fmt"
)

const (
	subOneMask uint64 = 0x0101010101010101
	matchAnd   uint64 = 0x8080808080808080
)

func FindPattern(word uint64, pattern uint64) int {
	dump("word^pattern", word, pattern, word^pattern)
	var match = word ^ pattern

	dump("match - matchMask", match, subOneMask, match-subOneMask)
	mask := match - subOneMask

	dump("mask & ^match", mask, ^match, mask & ^match)
	mask = mask & ^match

	dump("mask & matchAnd", mask, matchAnd, mask&matchAnd)
	mask = mask & matchAnd

	trailing := countTrailingZeros2(mask)
	return trailing >> 3
}

//func countTrailingZeros1[I integer](v I) (c int) {
//	fmt.Printf("trailing N %s\n", spaced(v))
//	for c = 0; v > 0; c++ {
//		dump("v & v-1", v, v-1, v&v-1)
//		v &= v - 1 // clear the least significant bit set
//	}
//	fmt.Printf("count %d\n", c)
//	return
//}

func countTrailingZeros2[I integer](v I) int {
	var c = 32

	v &= I(-int64(v))
	if v > 0 {
		c--
	}
	if v&0x0000FFFF > 0 {
		c -= 16
	}
	if v&0x00FF00FF > 0 {
		c -= 8
	}
	if v&0x0F0F0F0F > 0 {
		c -= 4
	}
	if v&0x33333333 > 0 {
		c -= 2
	}
	if v&0x55555555 > 0 {
		c -= 1
	}
	return c
}

type integer interface {
	~int64 | ~uint64 | int
}

func dump[I integer](s string, a, b, f I) {
	fmt.Printf("[%s]\n\t%s\n\t%s\n\t%s\n", s, spaced(a), spaced(b), spaced(f))
}

type spaced uint64

func (v spaced) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		for i := 7; i >= 0; i-- {
			mask := uint64(0xff) << (i * 8)
			vv := (uint64(v) & mask) >> (i * 8)
			fmt.Fprintf(s, "%08b ", vv)
		}
	default:
		fmt.Fprintf(s, "%064b", uint64(v))
	}
}
