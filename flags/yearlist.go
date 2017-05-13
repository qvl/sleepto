package flags

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type yearlist struct {
	list []int
}

func (l *yearlist) String() string {
	s := make([]string, len(l.list))
	for i := range l.list {
		s[i] = strconv.Itoa(l.list[i])
	}
	return strings.Join(s, ",")
}

func (l *yearlist) Set(s string) error {
	parts := strings.Split(s, ",")
	for i, p := range parts {
		x, err := strconv.Atoi(p)
		if err != nil {
			return fmt.Errorf("no integer at index %d: %s", i, p)
		}
		l.list = append(l.list, x)
	}
	return nil
}

// Yearlist defines a flag for a comma-separated list of integers.
// Call the returned function after flag.Parse to get the value.
func Yearlist(name string) func() []int {
	l := &yearlist{}
	flag.Var(l, name, "list of years")
	return func() []int {
		return l.list
	}
}
