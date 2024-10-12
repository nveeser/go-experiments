package puzzle

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type action string

const (
	switchOn   = "SwitchOn"
	switchOff  = "SwitchOff"
	switchNoOp = "NoOp"
)

var debug = false

func Execute(people, maxRuns int) (count int, found bool) {
	persons := makeRoster(people)
	var light bulb = false
	if debug {
		fmt.Printf("[%d] %s (%s)\n", count, light, dumpCounts(persons))
	}
	for ; count < maxRuns; count++ {
		pn := rand.IntN(len(persons))
		p := persons[pn]
		action := p.flip(light)

		switch action {
		case switchOff:
			light = false
		case switchOn:
			light = true
		}
		if debug {
			fmt.Printf("[%d] %s (%s) - %s %s\n", count+1, light, dumpCounts(persons), p, action)
		}

		if persons[pn].complete() {
			return count, true
		}
	}
	return count, false
}

func makeRoster(total int) []*person {
	persons := make([]*person, total)
	for i := 0; i < total; i++ {
		persons[i] = &person{
			id:     i,
			tokens: 1,
			total:  total,
		}
	}
	return persons
}

type person struct {
	id     int
	tokens int
	total  int
}

func (p *person) flip(light bulb) action {
	switch {
	case p.tokens == 0:
		return switchNoOp

	case light == true:
		p.tokens++
		return switchOff

	case light == false:
		p.tokens--
		return switchOn

	default:
		panic("What???")
	}
}
func (p *person) String() string {
	return fmt.Sprintf("Person[%d] (%d/%d)", p.id, p.tokens, p.total)
}

func (p *person) complete() bool {
	return p.tokens == p.total
}

type bulb bool

func (b bulb) String() string {
	if b {
		return "On"
	}
	return "Off"
}

func dumpCounts(p []*person) string {
	var s []string
	for _, pp := range p {
		s = append(s, fmt.Sprintf("%d", pp.tokens))
	}
	return strings.Join(s, ", ")
}
