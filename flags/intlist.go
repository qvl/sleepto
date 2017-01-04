package flags

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type intlist struct {
	min  int
	max  int
	list []int
}

func (l *intlist) String() string {
	s := make([]string, len(l.list))
	for i := range l.list {
		s[i] = strconv.Itoa(l.list[i])
	}
	return strings.Join(s, ",")
}

func (l *intlist) Set(s string) error {
	parts := strings.Split(s, ",")
	for i, p := range parts {
		x, err := strconv.Atoi(p)
		if err != nil {
			return fmt.Errorf("no integer at index %d: %s", i, p)
		}
		if x < l.min {
			return fmt.Errorf("%d is smaller than minimum %d", x, l.min)
		}
		if x > l.max {
			return fmt.Errorf("%d is bigger than maximum %d", x, l.max)
		}
		l.list = append(l.list, x)
	}
	return nil
}

// Intlist defines a flag for a comma-separated list of integers.
func Intlist(name, usage string, min, max int) []int {
	l := &intlist{min: min, max: max}
	flag.Var(l, name, usage)
	return l.list
}
