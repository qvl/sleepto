package flags

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type monthlist struct {
	list []time.Month
}

func (l *monthlist) String() string {
	s := make([]string, len(l.list))
	for i := range l.list {
		s[i] = strconv.Itoa(int(l.list[i]))
	}
	return strings.Join(s, ",")
}

func (l *monthlist) Set(s string) error {
	parts := strings.Split(s, ",")
	for i, p := range parts {
		xint, err := strconv.Atoi(p)
		if err != nil {
			return fmt.Errorf("no integer at index %d: %s", i, p)
		}
		x := time.Month(xint)
		if x < 1 || x > 12 {
			return fmt.Errorf("invalid month: %d", x)
		}
		l.list = append(l.list, x)
	}
	return nil
}

// Monthlist defines a flag for a comma-separated list of months.
// Valid values are between 1 and 12.
// Call the returned function after flag.Parse to get the value.
func Monthlist(name string) func() []time.Month {
	l := &monthlist{}
	flag.Var(l, name, "1 to 12")
	return func() []time.Month {
		return l.list
	}
}
