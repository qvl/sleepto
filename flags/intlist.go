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
		s[i] = string(l.list[i])
	}
	return strings.Join(s, ",")
}

func (l *intlist) Set(s string) error {
	parts := strings.Split(s, ",")
	for i, p := range parts {
		i64, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return fmt.Errorf("no integer at index %d: %s", i, p)
		}
		x := int(i64)
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

// Intlist ...
func Intlist(name, usage string, min, max int) []int {
	l := &intlist{min: min, max: max}
	flag.Var(l, name, usage)
	return l.list
}
