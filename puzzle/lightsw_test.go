package puzzle

import (
	"fmt"
	"testing"
)

func TestExecute(t *testing.T) {
	const max = 10000000
	for _, people := range []int{3, 10, 100} {
		t.Run(fmt.Sprintf("%d", people), func(t *testing.T) {
			count, found := Execute(people, max)
			t.Logf("Done: Total=%d Found=%t\n", count, found)
			if !found {
				t.Errorf("Run(len=%d, %d) did not complete", people, max)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	const max = 10000000
	for _, people := range []int{3, 10, 100} {
		b.Run(fmt.Sprintf("%d", people), func(b *testing.B) {
			var total float64
			for i := 0; i < b.N; i++ {
				count, _ := Execute(people, max)
				total += float64(count)
			}
			//b.ReportMetric(0, "ns/op")
			b.ReportMetric(total/float64(b.N), "iterations/op")
		})
	}
}
