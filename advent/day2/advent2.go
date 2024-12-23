package day2

import (
	"cmp"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Report []int

func (r Report) IsSafe() bool {
	var dir int
	for _, v := range r.fold() {
		if dir == 0 {
			dir = v.compare()
		}
		if !v.safe() || dir != v.compare() {
			return false
		}
	}
	return true
}

func (r Report) IsSafeDamped() bool {
	if r.IsSafe() {
		return true
	}
	for i := range r {
		r2 := slices.Delete(slices.Clone(r), i, i+1)
		if r2.IsSafe() {
			return true
		}
	}
	return false
}

func (r Report) fold() []*delta {
	var deltas []*delta
	for i := range r {
		if i+1 == len(r) {
			continue
		}
		deltas = append(deltas, &delta{r[i], r[i+1]})
	}
	return deltas
}

type delta struct{ n1, n2 int }

func (d *delta) String() string { return fmt.Sprintf("%d", d.n2-d.n1) }
func (d *delta) compare() int   { return cmp.Compare(d.n1, d.n2) }
func (d *delta) safe() bool {
	diff := d.n2 - d.n1
	if diff < 0 {
		diff = -diff
	}
	return 1 <= diff && diff <= 3
}

func ParseReports(data string) (reports []Report) {
	for _, row := range strings.Split(data, "\n") {
		row = strings.TrimFunc(row, unicode.IsSpace)
		if row == "" {
			continue
		}
		var report []int
		for _, field := range strings.Fields(row) {
			n, err := strconv.Atoi(field)
			if err != nil {
				log.Fatalf("error parsing: %s", err)
			}
			report = append(report, n)
		}
		reports = append(reports, report)
	}
	return
}
