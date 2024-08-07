package bitbit

import (
	"encoding/hex"
	"fmt"
	"testing"
	"unsafe"
)

const pattern uint64 = 0x3B3B3B3B3B3B3B3B

func TestFindPattern(t *testing.T) {
	a := []byte("skdjv;s\n")
	//a := []byte{0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7}
	fmt.Printf("Data: %s\n", hex.EncodeToString(a))
	ptr := unsafe.Pointer(&a[0])
	iptr := (*uint64)(unsafe.Pointer(uintptr(ptr)))
	val := *iptr
	fmt.Printf("Data: %016x\n", val)

	index := FindPattern(val, pattern)
	index = 8 - index
	fmt.Printf("Index: %d\n", index)
	if index != 5 {
		t.Errorf("FindPattern() got %d wanted %d", index, 5)
	}
}

func TestCountTrailing(t *testing.T) {
	cases := []struct {
		name  string
		input int64
		want  int
	}{
		{
			input: 0xF0,
			want:  4,
		},
		{
			input: 0x80,
			want:  7,
		},
		{
			input: 0x43,
			want:  0,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%08b", tc.input), func(t *testing.T) {
			got := countTrailingZeros2(tc.input)
			if got != tc.want {
				t.Errorf("count() got %d wanted %d", got, tc.want)
			}
		})
	}
}
