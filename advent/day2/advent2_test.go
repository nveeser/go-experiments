package day2

import (
	"fmt"
	"os"
	"testing"
)

func TestReports(t *testing.T) {
	cases := []struct {
		report   Report
		want     bool
		wantDrop bool
	}{
		// Increasing
		{[]int{1, 2, 3, 4}, true, true},  // Monotonic UUU
		{[]int{3, 2, 3, 4}, false, true}, // Monotonic DUU
		{[]int{1, 4, 3, 6}, false, true}, // Monotonic UDU
		{[]int{1, 2, 3, 2}, false, true}, // Monotonic UUD
		// Decreasing
		{[]int{5, 4, 3, 2}, true, true},  // Monotonic DDD
		{[]int{1, 4, 3, 2}, false, true}, // Monotonic UDD
		{[]int{5, 3, 4, 2}, false, true}, // Monotonic DUD
		{[]int{5, 3, 2, 3}, false, true}, // Monotonic DDU
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Logf("Report %+v", tc.report)
			t.Run("isSafe()", func(t *testing.T) {
				got := tc.report.IsSafe()
				if got != tc.want {
					t.Errorf("got %t wanted %t", got, tc.want)
				}
			})
			t.Run("IsSafeDamped()", func(t *testing.T) {
				got := tc.report.IsSafeDamped()
				if got != tc.wantDrop {
					t.Errorf("got %t wanted %t", got, tc.wantDrop)
				}
			})
		})
	}
}

func TestCount(t *testing.T) {
	small := ParseReports(smallData)
	large := ParseReports(largeData(t))
	t.Run("Safe/Small", func(t *testing.T) {
		got := countSafe(t, small)
		if got != 2 {
			t.Errorf("countSafe got %d wanted %d", got, 2)
		}
	})
	t.Run("Safe/Large", func(t *testing.T) {
		got := countSafe(t, large)
		if got != 257 {
			t.Errorf("countSafe got %d wanted %d", got, 257)
		}
	})
	t.Run("SafeDamped/Small", func(t *testing.T) {
		got := countSafeDamped(t, small)
		if got != 5 {
			t.Errorf("CheckReports got %d wanted %d", got, 5)
		}
	})
	t.Run("SafeDamped/Large", func(t *testing.T) {
		got := countSafeDamped(t, large)
		if got != 328 {
			t.Errorf("CheckReports got %d wanted %d", got, 328)
		}
	})
}

func countSafe(t *testing.T, reports []Report) int {
	var safe int
	for _, r := range reports {
		t.Logf("Report: %+v => %+v => %t\n", r, r.fold(), r.IsSafe())
		if r.IsSafe() {
			safe++
		}
	}
	return safe
}

func countSafeDamped(t *testing.T, reports []Report) int {
	var safe int
	for _, r := range reports {
		t.Logf("Report: %+v => %+v => %t\n", r, r.fold(), r.IsSafe())
		if r.IsSafeDamped() {
			safe++
		}
	}
	return safe
}

const smallData = `
3 1 3 5 7
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func largeData(t *testing.T) string {
	data, err := os.ReadFile("large.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %s", "large.txt")
	}
	return string(data)
}
